/*
 * @lc app=leetcode.cn id=34 lang=golang
 * @lcpr version=21909
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package main

func lowerBound3(nums []int, target int) int {
	left, right := -1, len(nums) // 开区间 (left, right)
	for left+1 < right {         // 区间不为空
		// 循环不变量：
		// nums[left] < target
		// nums[right] >= target
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid // 范围缩小到 (mid, right)
		} else {
			left = mid // 范围缩小到 (left, mid)
		}
	}
	return right // 或者 left+1
}
func searchRange(nums []int, target int) []int {
	start := lowerBound3(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}

	end := lowerBound3(nums, target+1) - 1
	return []int{start, end}
}

// @lc code=end

/*
// @lcpr case=start
// [5,7,7,8,8,10]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,7,7,8,8,10]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/
