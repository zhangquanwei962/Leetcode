/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=21909
 *
 * [15] 三数之和
 */

// @lc code=start
package main

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] { // 跳过重复数字
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // 优化一
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 { // 优化二
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过重复数字
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/
