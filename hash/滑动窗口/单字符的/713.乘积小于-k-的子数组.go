/*
 * @lc app=leetcode.cn id=713 lang=golang
 * @lcpr version=21909
 *
 * [713] 乘积小于 K 的子数组
 */

// @lc code=start
package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	ans, left := 0, 0
	prod := 1

	for right, x := range nums {
		prod *= x
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,2,6]\n100\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n0\n
// @lcpr case=end

*/
