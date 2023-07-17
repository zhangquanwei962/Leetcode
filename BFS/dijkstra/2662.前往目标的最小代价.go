/*
 * @lc app=leetcode.cn id=2662 lang=golang
 * @lcpr version=21909
 *
 * [2662] 前往目标的最小代价
 */

// @lc code=start
// O(n2) O(n)
package main

import "math"

func minimumCost(start, target []int, specialRoads [][]int) int {
	type pair struct{ x, y int }
	t := pair{target[0], target[1]}
	dis := make(map[pair]int, len(specialRoads)+2)
	dis[t] = math.MaxInt
	dis[pair{start[0], start[1]}] = 0
	vis := make(map[pair]bool, len(specialRoads)+1) // 终点不用记
	for {
		v, dv := pair{}, -1 // 类似之前
		for p, d := range dis {
			if !vis[p] && (v == pair{} || d < dv) { // dv<0
				v, dv = p, d
			}
		}
		if v == t { // 到终点的最短路已确定
			return dv
		}
		vis[v] = true
		dis[t] = min(dis[t], dv+t.x-v.x+t.y-v.y) // 更新到终点的最短路
		for _, r := range specialRoads {
			w := pair{r[2], r[3]} // 要么直接走曼哈顿，要么先到r[0] r[1]再走特殊路径
			d := dv + abs(r[0]-v.x) + abs(r[1]-v.y) + r[4]
			if dw, ok := dis[w]; !ok || d < dw {
				dis[w] = d
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,1]\n[4,5]\n[[1,2,3,3,2],[3,4,4,5,1]]\n
// @lcpr case=end

// @lcpr case=start
// [3,2]\n[5,7]\n[[3,2,3,4,4],[3,3,5,5,5],[3,4,5,6,6]]\n
// @lcpr case=end

*/
