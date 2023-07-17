/*
 * @lc app=leetcode.cn id=2465 lang=golang
 * @lcpr version=21908
 *
 * [2465] 不同的平均值数目
 */

// @lc code=start
package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	s := map[int]struct{}{}

	for i := 0; i < n>>1; i++ {
		s[nums[i]+nums[n-i-1]] = struct{}{}
	}
	return len(s)
}

// @lc code=end

/*
// @lcpr case=start
// [4,1,4,0,3,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,100]\n
// @lcpr case=end

*/
