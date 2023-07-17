/*
 * @lc app=leetcode.cn id=2101 lang=golang
 * @lcpr version=21909
 *
 * [2101] 引爆最多的炸弹
 */

// @lc code=start
// 暴力枚举dfs，看能跑多少
package main

func maximumDetonation(bombs [][]int) (ans int) {
	n := len(bombs)
	g := make([][]int, n) // 建图
	for i, p := range bombs {
		for j, q := range bombs {
			if i != j && (q[0]-p[0])*(q[0]-p[0])+(q[1]-p[1])*(q[1]-p[1]) <= p[2]*p[2] {
				g[i] = append(g[i], j)
			}
		}
	}

	for i := range g {
		vis := make([]bool, n)
		cnt := 0
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			cnt++
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
		}
		dfs(i)
		if cnt > ans {
			ans = cnt
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,3],[6,1,4]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,5],[10,10,5]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]\n
// @lcpr case=end

*/
