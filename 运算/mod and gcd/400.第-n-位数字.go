/*
 * @lc app=leetcode.cn id=400 lang=golang
 * @lcpr version=21909
 *
 * [400] 第 N 位数字
 */

// @lc code=start
// 定位到段：确定该位位于第几段，即确定 k 。
// 定位到段内的数字：确定该位位于第 k 段的第几个数，即 num 。
// 定位该数字的偏移量：确定该位位于 num 的第几位，记作 offset 。
package main

func findNthDigit(n int) int {
	d := 1 // 确定几位数
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}
	// one-base count
	index := n - 1
	start := qmi(10, d-1)
	num := start + index/d
	digitIndex := index % d
	return num / qmi(10, d-digitIndex-1) % 10
}

// O(logn)
func qmi(a, b int) int {
	res := 1
	for ; b > 0; b >>= 1 {
		if b&1 == 1 {
			res = res * a
		}
		a = a * a
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 11\n
// @lcpr case=end

*/
