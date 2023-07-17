/*
 * @lc app=leetcode.cn id=231 lang=golang
 * @lcpr version=21909
 *
 * [231] 2 的幂
 */

// @lc code=start
package main

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0 && n > 0
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 16\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 5\n
// @lcpr case=end

*/
