/*
 * @lc app=leetcode.cn id=852 lang=golang
 * @lcpr version=21909
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
// 找到第一个>=的
package main

func peakIndexInMountainArray(arr []int) int {
	left, right := -1, len(arr)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if arr[mid] < arr[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

// func findPeakElement(nums []int) int {
// 	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
// 	for left+1 < right {           // 开区间不为空
// 		mid := left + (right-left)>>1
// 		if nums[mid] > nums[mid+1] { // 蓝色
// 			right = mid
// 		} else { // 红色
// 			left = mid
// 		}
// 	}
// 	return right
// }

// @lc code=end

/*
// @lcpr case=start
// [0,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,2,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,10,5,2]\n
// @lcpr case=end

*/
