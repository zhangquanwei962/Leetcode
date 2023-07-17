/*
 * @lc app=leetcode.cn id=162 lang=golang
 * @lcpr version=21909
 *
 * [162] 寻找峰值
 */

// @lc code=start
// 找到第一个 nums[i] > nums[i+1] 的位置即可
package main

func findPeakElement(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid+1] { // 蓝色
			right = mid
		} else { // 红色
			left = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,3,5,6,4]\n
// @lcpr case=end

*/
