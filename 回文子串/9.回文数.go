/*
 * @lc app=leetcode.cn id=9 lang=golang
 * @lcpr version=21909
 *
 * [9] 回文数
 */

// @lc code=start
package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	// y := strconv.Itoa(x)
	y := strconv.FormatInt(int64(x), 10)

	n := len(y)

	for i := 0; i < n>>1; i++ {
		if y[i] != y[n-i-1] {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// 121\n
// @lcpr case=end

// @lcpr case=start
// -121\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/
