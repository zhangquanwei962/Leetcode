/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=21909
 *
 * [53] 最大子数组和
 */

// @lc code=start
package main

import "math"

// func maxSubArray(nums []int) int {
// 	res, pre := nums[0], nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		pre = max(nums[i], pre+nums[i])
// 		res = max(pre, res)
// 	}
// 	return res
// }

// func maxSubArray(nums []int) int {
// 	maxSum, s := math.MinInt/2, 0
// 	for _, x := range nums {
// 		s = max(s, 0) + x
// 		maxSum = max(maxSum, s)
// 	}
// 	return maxSum
// }

func maxSubArray(nums []int) int {
	maxSum := math.MinInt
	dp := make([]int, len(nums)+1)
	dp[0] = math.MinInt
	for i, x := range nums {
		dp[i+1] = max(dp[i], 0) + x
		maxSum = max(dp[i+1], maxSum)
	}
	return maxSum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// package main

// import "math"

// func maxSubArray(arr []int) int {
// 	memo := make([]int, len(arr))
// 	for i := range memo {
// 		memo[i] = math.MinInt
// 	}

// 	var dfs func(int) int
// 	dfs = func(i int) (res int) {
// 		if i < 0 {
// 			return math.MinInt / 2
// 		}
// 		p := &memo[i]
// 		if *p != math.MinInt { // 之前计算过
// 			return *p
// 		}
// 		defer func() { *p = res }() // 记忆化
// 		return max(dfs(i-1), 0) + arr[i]

// 	}
// 	ans := math.MinInt
// 	for i := range arr {
// 		ans = max(ans, dfs(i))
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if b > a {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/
