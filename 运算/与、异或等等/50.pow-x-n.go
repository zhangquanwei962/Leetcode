/*
 * @lc app=leetcode.cn id=50 lang=golang
 * @lcpr version=21909
 *
 * [50] Pow(x, n)
 */

// @lc code=start
package main

func qmi(x float64, k int) float64 {
	res := 1.0
	for ; k != 0; k >>= 1 {
		if k&1 > 0 {
			res *= x
		}
		x *= x
	}
	return res
}
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return qmi(x, n)
	}
	return 1.0 / qmi(x, -n)
}

// @lc code=end

/*
// @lcpr case=start
// 2.00000\n10\n
// @lcpr case=end

// @lcpr case=start
// 2.10000\n3\n
// @lcpr case=end

// @lcpr case=start
// 2.00000\n-2\n
// @lcpr case=end

*/
