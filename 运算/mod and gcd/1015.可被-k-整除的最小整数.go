/*
 * @lc app=leetcode.cn id=1015 lang=golang
 * @lcpr version=21908
 *
 * [1015] 可被 K 整除的最小整数
 */

// @lc code=start
package main

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	n := 1 % k
	for i := 1; i <= k; i++ {
		if n == 0 {
			return i
		}
		n = (n*10 + 1) % k
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/
