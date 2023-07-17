/*
 * @lc app=leetcode.cn id=1911 lang=golang
 * @lcpr version=21908
 *
 * [1911] 最大子序列交替和
 */

// @lc code=start
package main

import "math"

func maxAlternatingSum(nums []int) int64 {
	f := [2]int{0, math.MinInt64 / 2}

	for _, v := range nums {
		f = [2]int{max(f[0], f[1]-v), max(f[1], f[0]+v)}
	}

	return int64(f[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,5,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,6,7,8]\n
// @lcpr case=end

// @lcpr case=start
// [6,2,1,2,4,5]\n
// @lcpr case=end

*/
