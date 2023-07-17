/*
 * @lc app=leetcode.cn id=1201 lang=golang
 * @lcpr version=21909
 *
 * [1201] 丑数 III
 */

// @lc code=start
// <=x的神奇数字的有n个
package main

func nthUglyNumber(n int, a int, b int, c int) int {
	lcm1 := a / gcd(a, b) * b
	lcm2 := a / gcd(a, c) * c
	lcm3 := b / gcd(b, c) * c
	lcm4 := lcm1 / gcd(lcm1, c) * c
	left, right := 0, min(c, min(a, b))*n
	for left+1 < right {
		mid := left + (right-left)>>1
		if mid/a+mid/b+mid/c-mid/lcm1-mid/lcm2-mid/lcm3+mid/lcm4 >= n {
			right = mid
		} else {
			left = mid
		}
	}
	return right
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
// 3\n2\n3\n5\n
// @lcpr case=end

// @lcpr case=start
// 4\n2\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// 5\n2\n11\n13\n
// @lcpr case=end

// @lcpr case=start
// 1000000000\n2\n217983653\n336916467\n
// @lcpr case=end

*/
