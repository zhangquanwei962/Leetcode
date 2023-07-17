/*
 * @lc app=leetcode.cn id=448 lang=golang
 * @lcpr version=21909
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
package main

func findDisappearedNumbers(nums []int) (res []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return

}

// @lc code=end

/*
// @lcpr case=start
// [4,3,2,7,8,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/
