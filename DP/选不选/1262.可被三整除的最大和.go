/*
 * @lc app=leetcode.cn id=1262 lang=golang
 * @lcpr version=21909
 *
 * [1262] 可被三整除的最大和
 */

// @lc code=start
package main

import "math"

func maxSumDivThree(nums []int) int {
	f := [2][3]int{{0, math.MinInt, math.MinInt}}
	for i, x := range nums {
		for j := 0; j < 3; j++ {
			f[(i+1)%2][j] = max(f[i%2][j], f[i%2][(j+x)%3]+x)
		}
	}
	return f[len(nums)%2][0]
}

// func maxSumDivThree(nums []int) int {
// 	n := len(nums)
// 	f := make([][3]int, n+1)
// 	f[0][1] = math.MinInt
// 	f[0][2] = math.MinInt

// 	for i, x := range nums {
// 		for j := 0; j < 3; j++ {
// 			f[i+1][j] = max(f[i][j], f[i][(j+x)%3]+x)
// 		}
// 	}
// 	return f[n][0]
// }

// func maxSumDivThree(nums []int) int {
// 	n := len(nums)
// 	memo := make([][3]int, n)
// 	for i := range memo {
// 		for j := range memo[i] {
// 			memo[i][j] = -1
// 		}
// 	}

// 	var dfs func(int, int) int
// 	dfs = func(i, j int) int {
// 		if i < 0 {
// 			if j == 0 {
// 				return 0
// 			}
// 			return math.MinInt
// 		}

// 		p := &memo[i][j]
// 		if *p != -1 {
// 			return *p
// 		}
// 		*p = max(dfs(i-1, j), dfs(i-1, (j+nums[i])%3)+nums[i])
// 		return *p
// 	}
// 	return dfs(n-1, 0)
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [3,6,5,1,8]\n
// @lcpr case=end

// @lcpr case=start
// [4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,4]\n
// @lcpr case=end

*/
