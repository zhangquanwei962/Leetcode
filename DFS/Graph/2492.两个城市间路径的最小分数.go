/*
 * @lc app=leetcode.cn id=2492 lang=golang
 * @lcpr version=21909
 *
 * [2492] 两个城市间路径的最小分数
 */

// @lc code=start
package main

import "math"

func minScore(n int, roads [][]int) int {
	type edge struct{ to, d int }
	g := make([][]edge, n)
	for _, e := range roads {
		x, y, d := e[0]-1, e[1]-1, e[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}

	ans := math.MaxInt32
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, e := range g[x] {
			ans = min(ans, e.d)
			if !vis[e.to] {
				dfs(e.to)
			}
		}
	}
	dfs(0)
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// 4\n[[1,2,9],[2,3,6],[2,4,5],[1,4,7]]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[1,2,2],[1,3,4],[3,4,7]]\n
// @lcpr case=end

*/
