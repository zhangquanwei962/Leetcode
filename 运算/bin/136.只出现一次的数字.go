/*
 * @lc app=leetcode.cn id=136 lang=golang
 * @lcpr version=21909
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
package main

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,1,2,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
