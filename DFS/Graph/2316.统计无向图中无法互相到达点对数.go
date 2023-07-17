/*
 * @lc app=leetcode.cn id=2316 lang=golang
 * @lcpr version=21908
 *
 * [2316] 统计无向图中无法互相到达点对数
 */

// @lc code=start
package main

func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	tot, size := 0, 0
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		size++
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if !b {
			size = 0
			dfs(i)
			ans += int64(size) * int64(tot)
			tot += size
		}
	}
	return

}

// @lc code=end

/*
// @lcpr case=start
// 3\n[[0,1],[0,2],[1,2]]\n
// @lcpr case=end

// @lcpr case=start
// 7\n[[0,2],[0,5],[2,4],[1,6],[5,4]]\n
// @lcpr case=end

*/
