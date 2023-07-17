/*
 * @lc app=leetcode.cn id=168 lang=golang
 * @lcpr version=21909
 *
 * [168] Excel表列名称
 */

// @lc code=start
// 因为从1开始，需要先减1
package main

func convertToTitle(columnNumber int) (ans string) {
	for columnNumber > 0 {
		columnNumber--
		ans = string(rune(columnNumber%26+'A')) + ans
		columnNumber /= 26
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 28\n
// @lcpr case=end

// @lcpr case=start
// 701\n
// @lcpr case=end

// @lcpr case=start
// 2147483647\n
// @lcpr case=end

*/
