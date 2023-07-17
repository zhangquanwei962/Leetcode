/*
 * @lc app=leetcode.cn id=62 lang=golang
 * @lcpr version=21907
 *
 * [62] 不同路径
 */

// @lc code=start
// 网格图机器人一定会走m+n-2步
package main

func uniquePaths(m int, n int) int {
	// dp := make([][]int, m)
	// for i := range dp {
	// 	dp[i] = make([]int, n)
	// 	dp[i][0] = 1
	// }
	// for j := 0; j < n; j++ {
	// 	dp[0][j] = 1
	// }

	// for i := 1; i < m; i++ {
	// 	for j := 1; j < n; j++ {
	// 		dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// 	}
	// }
	// return dp[m-1][n-1]

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没被计算过
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return
		}
		if i == 0 && j == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		res += dfs(i-1, j) + dfs(i, j-1)
		return res
	}
	return dfs(m-1, n-1)
}

// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n2\n
// @lcpr case=end

// @lcpr case=start
// 7\n3\n
// @lcpr case=end

// @lcpr case=start
// 3\n3\n
// @lcpr case=end

*/
