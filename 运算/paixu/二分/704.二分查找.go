/*
 * @lc app=leetcode.cn id=704 lang=golang
 * @lcpr version=21909
 *
 * [704] 二分查找
 */

// @lc code=start
package main

func search(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	if right == len(nums) || nums[right] != target {
		return -1
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,3,5,9,12]\n9\n
// @lcpr case=end

// @lcpr case=start
// [-1,0,3,5,9,12]\n2\n
// @lcpr case=end

*/
