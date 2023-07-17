/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=21909
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
// O(n) O(1)
// 对于环形，非环形正常求，环形求中间最小和
package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	maxSum, minSum := math.MinInt, math.MaxInt
	maxS, minS, arrSum := 0, 0, 0

	for _, x := range nums {
		arrSum += x
		maxS = max(maxS, 0) + x
		minS = min(minS, 0) + x
		maxSum = max(maxSum, maxS)
		minSum = min(minSum, minS)
	}

	if maxSum < 0 {
		return maxSum
	}
	return max(arrSum-minSum, maxSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
// [1,-2,3,-2]\n
// @lcpr case=end

// @lcpr case=start
// [5,-3,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2,2,-3]\n
// @lcpr case=end

*/
