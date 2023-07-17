/*
 * @lc app=leetcode.cn id=201 lang=golang
 * @lcpr version=21909
 *
 * [201] 数字范围按位与
 */

// @lc code=start
// &运算，得到的结果就是小于等于俩得最小值
// 前缀相等
package main

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= right - 1
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// 5\n7\n
// @lcpr case=end

// @lcpr case=start
// 0\n0\n
// @lcpr case=end

// @lcpr case=start
// 1\n2147483647\n
// @lcpr case=end

*/
