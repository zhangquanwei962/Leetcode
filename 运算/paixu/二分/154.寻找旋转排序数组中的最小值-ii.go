/*
 * @lc app=leetcode.cn id=154 lang=golang
 * @lcpr version=21909
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
package main

func findMin(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)/2
		if nums[mid] < nums[right] { // 蓝色
			right = mid
		} else if nums[mid] > nums[right] { // 红色
			left = mid
		} else {
			right--
		}
	}
	return nums[right]
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,5]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,0,1]\n
// @lcpr case=end

*/
