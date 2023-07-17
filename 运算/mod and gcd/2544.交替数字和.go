/*
 * @lc app=leetcode.cn id=2544 lang=golang
 * @lcpr version=21909
 *
 * [2544] 交替数字和
 */

// @lc code=start
// sign 符号
package main

func alternateDigitSum(n int) (ans int) {
	sign := 1
	for n > 0 {
		ans += n % 10 * sign
		n /= 10
		sign = -sign
	}
	return ans * (-sign)
}

// @lc code=end

/*
// @lcpr case=start
// 521\n
// @lcpr case=end

// @lcpr case=start
// 111\n
// @lcpr case=end

// @lcpr case=start
// 886996\n
// @lcpr case=end

*/
