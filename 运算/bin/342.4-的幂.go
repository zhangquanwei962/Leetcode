/*
 * @lc app=leetcode.cn id=342 lang=golang
 * @lcpr version=21907
 *
 * [342] 4的幂
 */

// @lc code=start
package main

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}

// @lc code=end

/*
// @lcpr case=start
// 16\n
// @lcpr case=end

// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
