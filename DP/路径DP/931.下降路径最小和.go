/*
 * @lc app=leetcode.cn id=931 lang=golang
 * @lcpr version=21908
 *
 * [931] 下降路径最小和
 */

// @lc code=start
package main

import (
	"math"
)

func minFallingPathSum(A [][]int) int {
	// 设dp[i][j]为到i, j位置的最小路径和
	len := len(A)
	dp := make([][]int, len+1)
	for i := range dp {
		dp[i] = make([]int, len+2)
	}

	// 套壳处理
	for i := 0; i < len+1; i++ {
		dp[i][0] = math.MaxInt
		dp[i][len+1] = math.MaxInt
	}
	for j := 0; j < len+2; j++ {
		dp[0][j] = 0
	}

	ans := math.MaxInt
	for i := 1; i < len+1; i++ {
		for j := 1; j < len+1; j++ {
			dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i-1][j+1]) + A[i-1][j-1]
		}
	}
	for i := 1; i < len+1; i++ {
		ans = min(ans, dp[len][i])
	}
	return ans
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
// [[2,1,3],[6,5,4],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[-19,57],[-40,-5]]\n
// @lcpr case=end

*/
