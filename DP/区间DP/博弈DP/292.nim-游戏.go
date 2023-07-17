/*
 * @lc app=leetcode.cn id=292 lang=golang
 * @lcpr version=21909
 *
 * [292] Nim 游戏
 */

// @lc code=start
package main

func canWinNim(n int) bool {
	return n%4 != 0
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/
