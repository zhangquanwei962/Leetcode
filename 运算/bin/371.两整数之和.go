/*
 * @lc app=leetcode.cn id=371 lang=golang
 * @lcpr version=21909
 *
 * [371] 两整数之和
 */

// @lc code=start
// 无进位加法 a^b
// 有进位 a&b<<1
package main

func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// 1\n2\n
// @lcpr case=end

// @lcpr case=start
// 2\n3\n
// @lcpr case=end

*/
