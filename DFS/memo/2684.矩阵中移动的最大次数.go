/*
 * @lc app=leetcode.cn id=2684 lang=golang
 * @lcpr version=21908
 *
 * [2684] 矩阵中移动的最大次数
 */

// @lc code=start
package main

func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j == n-1 {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		for k := max(i-1, 0); k < min(i+2, m); k++ {
			if grid[k][j+1] > grid[i][j] {
				res = max(res, dfs(k, j+1)+1)
			}
		}
		return
	}
	for i := 0; i < m; i++ {
		ans = max(ans, dfs(i, 0))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
// [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,2,4],[2,1,9],[1,1,7]]\n
// @lcpr case=end

*/
