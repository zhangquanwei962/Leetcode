/*
 * @lc app=leetcode.cn id=33 lang=golang
 * @lcpr version=21909
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package main

func findMin(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] { // 蓝色
			right = mid
		} else { // 红色
			left = mid
		}
	}
	return right
}

func lowerBound(nums []int, left, right, target int) int {
	r0 := right
	for left+1 < right { // 开区间不为空
		// 循环不变量：
		// nums[left] < target
		// nums[right] >= target
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid // 范围缩小到 (mid, right)
		} else {
			right = mid // 范围缩小到 (left, mid)
		}
	}
	if right == r0 || nums[right] != target {
		return -1
	}
	return right
}

func search(nums []int, target int) int {
	i := findMin(nums)
	if target > nums[len(nums)-1] {
		return lowerBound(nums, -1, i, target) // 左段
	}
	return lowerBound(nums, i-1, len(nums), target) // 右段
}

// @lc code=end

/*
// @lcpr case=start
// [4,5,6,7,0,1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/
