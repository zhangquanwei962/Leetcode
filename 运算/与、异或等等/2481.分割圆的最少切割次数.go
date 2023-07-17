/*
 * @lc app=leetcode.cn id=2481 lang=golang
 * @lcpr version=21909
 *
 * [2481] 分割圆的最少切割次数
 */

// @lc code=start
package main

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n&1 == 1 {
		return n
	}
	return n >> 1
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/
