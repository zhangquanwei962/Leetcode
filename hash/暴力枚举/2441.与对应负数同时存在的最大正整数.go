/*
 * @lc app=leetcode.cn id=2441 lang=golang
 * @lcpr version=21908
 *
 * [2441] 与对应负数同时存在的最大正整数
 */

// @lc code=start
package main

func findMaxK(nums []int) int {
	m := map[int]bool{}

	for _, x := range nums {
		m[x] = true
	}
	ans := -1

	for _, x := range nums {
		if m[-x] && ans < x {
			ans = x
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [-1,2,-3,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,10,6,7,-7,1]\n
// @lcpr case=end

// @lcpr case=start
// [-10,8,6,7,-2,-3]\n
// @lcpr case=end

*/
