/*
 * @lc app=leetcode.cn id=762 lang=golang
 * @lcpr version=21907
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
package main

import "math/bits"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
func countPrimeSetBits(left int, right int) (ans int) {
	for x := left; x <= right; x++ {
		if isPrime(bits.OnesCount(uint(x))) {
			ans++
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 6\n10\n
// @lcpr case=end

// @lcpr case=start
// 10\n15\n
// @lcpr case=end

*/
