/*
 * @lc app=leetcode.cn id=878 lang=golang
 * @lcpr version=21909
 *
 * [878] 第 N 个神奇数字
 */

// @lc code=start
// <=x的神奇数字的有n个
package main

func nthMagicalNumber(n int, a int, b int) int {
	lcm := a / gcd(a, b) * b
	left, right := 0, min(a, b)*n
	for left+1 < right {
		mid := left + (right-left)>>1
		if mid/a+mid/b-mid/lcm >= n {
			right = mid
		} else {
			left = mid
		}
	}
	return right % (1e9 + 7)
}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// 1\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 4\n2\n3\n
// @lcpr case=end

*/
