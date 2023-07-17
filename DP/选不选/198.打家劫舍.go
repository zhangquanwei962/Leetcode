/*
 * @lc app=leetcode.cn id=198 lang=golang
 * @lcpr version=21909
 *
 * [198] 打家劫舍
 */

// @lc code=start
package main

func rob(nums []int) int {
	f0, f1 := 0, 0
	for i := 0; i < len(nums); i++ {
		f0, f1 = f1, max(f1, f0+nums[i])
	}
	return f1
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
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

*/
