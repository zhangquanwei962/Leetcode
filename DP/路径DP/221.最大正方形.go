/*
 * @lc app=leetcode.cn id=221 lang=golang
 * @lcpr version=21909
 *
 * [221] 最大正方形
 */

// @lc code=start
package main

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	// dp表示以第i行，第j列为右下角的边长
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxSide := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == byte('1') {
				dp[i+1][j+1] = min(min(dp[i+1][j], dp[i][j+1]), dp[i][j]) + 1
				if dp[i+1][j+1] > maxSide {
					maxSide = dp[i+1][j+1]
				}
			}
		}
	}
	return maxSide * maxSide
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
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0","1"],["1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

*/
