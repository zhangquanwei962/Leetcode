/*
 * @lc app=leetcode.cn id=1269 lang=golang
 * @lcpr version=21909
 *
 * [1269] 停在原地的方案数
 */

// @lc code=start
// 和
package main

func numWays(steps, arrLen int) int {
	const mod = 1e9 + 7
	maxColumn := min(arrLen-1, steps)
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, maxColumn+1)
	}
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j <= maxColumn; j++ {
			dp[i][j] = dp[i-1][j] // 不动
			if j-1 >= 0 {         // 左移一次
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j+1 <= maxColumn { // 右移一次
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[steps][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// 3\n2\n
// @lcpr case=end

// @lcpr case=start
// 2\n4\n
// @lcpr case=end

// @lcpr case=start
// 4\n2\n
// @lcpr case=end

*/
