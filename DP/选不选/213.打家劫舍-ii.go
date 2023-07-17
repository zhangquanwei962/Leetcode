/*
 * @lc app=leetcode.cn id=213 lang=golang
 * @lcpr version=21909
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
package main

func rob1(nums []int, start, end int) int {
	f0, f1 := 0, 0
	for i := start; i < end; i++ {
		f0, f1 = f1, max(f1, f0+nums[i])
	}
	return f1
}

func rob(nums []int) int {
	n := len(nums)
	return max(nums[0]+rob1(nums, 2, n-1), rob1(nums, 1, n))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/
