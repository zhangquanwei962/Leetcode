/*
 * @lc app=leetcode.cn id=1289 lang=golang
 * @lcpr version=21909
 *
 * [1289] 下降路径最小和  II
 */

// @lc code=start
/*
可以这样写，在记忆化函数中传入的参数是 (0,n-1) ，其实这就是动态规划的出口，所以 i 最终要变成 0 ， j 最终要变成 n - 1，所以应该是 i 递减，j 递增
*/
package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt / 2
			for k := 0; k < n; k++ {
				if k != j {
					dp[i][j] = min(dp[i][j], dp[i-1][k]+matrix[i][j])
				}
			}
		}
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dp[m-1][i])
	}
	return res

}

// func minFallingPathSum(matrix [][]int) int {

// 	m, n := len(matrix), len(matrix[0])

// 	memo := make([][]int, m)
// 	for i := range memo {
// 		memo[i] = make([]int, n)
// 		for j := range memo[i] {
// 			memo[i][j] = -1 // -1 表示没被计算过
// 		}
// 	}

// 	var dfs func(int, int) int
// 	dfs = func(i, j int) int {
// 		if i == 0 {
// 			return matrix[0][j]
// 		}
// 		p := &memo[i][j]
// 		if *p != -1 {
// 			return *p
// 		}
// 		res := math.MaxInt
// 		for k := 0; k < n; k++ {
// 			if k != j {
// 				res = min(res, matrix[i][j]+dfs(i-1, k))
// 			}
// 		}
// 		defer func() { *p = res }()

// 		return res
// 	}
// 	res := math.MaxInt
// 	for i := 0; i < n; i++ {
// 		res = min(res, dfs(m-1, i))
// 	}
// 	return res

// }

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[7]]\n
// @lcpr case=end

*/
