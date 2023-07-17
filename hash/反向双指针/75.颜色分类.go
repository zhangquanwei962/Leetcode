/*
 * @lc app=leetcode.cn id=75 lang=golang
 * @lcpr version=21909
 *
 * [75] 颜色分类
 */

// @lc code=start
// 三路快排的思想
package main

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	i := 0
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			i++
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [2,0,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [2,0,1]\n
// @lcpr case=end

*/
