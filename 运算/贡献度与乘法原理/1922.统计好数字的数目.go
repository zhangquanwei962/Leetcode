/*
 * @lc app=leetcode.cn id=1922 lang=golang
 * @lcpr version=21909
 *
 * [1922] 统计好数字的数目
 */

// @lc code=start
// 快速幂 + 乘法原理
package main

const mod int = 1e9 + 7

func countGoodNumbers(n int64) int {
	return pow(5, (int(n)+1)/2) * pow(4, int(n)/2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 50\n
// @lcpr case=end

*/
