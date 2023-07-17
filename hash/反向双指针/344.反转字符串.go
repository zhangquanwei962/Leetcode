/*
 * @lc app=leetcode.cn id=344 lang=golang
 * @lcpr version=21909
 *
 * [344] 反转字符串
 */

// @lc code=start
package main

func reverseString(s []byte) {
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// ["h","e","l","l","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["H","a","n","n","a","h"]\n
// @lcpr case=end

*/
