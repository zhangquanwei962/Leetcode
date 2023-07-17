/*
 * @lc app=leetcode.cn id=153 lang=golang
 * @lcpr version=21909
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
package main

func findMin(nums []int) int {
	left, right := -1, len(nums)-1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right]
}

// @lc code=end

/*
// @lcpr case=start
// [3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [11,13,15,17]\n
// @lcpr case=end

*/
