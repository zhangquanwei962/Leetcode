/*
 * @lc app=leetcode.cn id=2467 lang=golang
 * @lcpr version=21908
 *
 * [2467] 树上最大得分和路径
 */

// @lc code=start
package main

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)

	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	bobTime := make([]int, n) // bobTime[x] 表示 bob 访问节点 x 的时间
	for i := range bobTime {
		bobTime[i] = n // 也可以初始化成 inf
	}

	var dfsBob func(int, int, int) bool
	dfsBob = func(x, fa, t int) bool {
		if x == 0 {
			bobTime[x] = t
			return true
		}

		for _, y := range g[x] {
			if y != fa && dfsBob(y, x, t+1) {
				bobTime[x] = t // 只有可以到达 0 才标记访问时间
				return true
			}
		}
		return false
	}

	dfsBob(bob, -1, 0)

	g[0] = append(g[0], -1) // 防止把根节点当作叶子
	ans := math.MinInt32
	var dfsAlice func(int, int, int, int)
	dfsAlice = func(x, fa, AilceTime, sum int) {
		if AilceTime < bobTime[x] {
			sum += amount[x]
		} else if AilceTime == bobTime[x] {
			sum += amount[x] / 2
		}

		if len(g[x]) == 1 { // 叶子
			ans = max(ans, sum) // 更新答案
			return
		}

		for _, y := range g[x] {
			if y != fa {
				dfsAlice(y, x, AilceTime+1, sum)
			}
		}
	}

	dfsAlice(0, -1, 0, 0)

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[1,2],[1,3],[3,4]]\n3\n[-2,4,2,-4,6]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1]]\n1\n[-7280,2350]\n
// @lcpr case=end

*/
