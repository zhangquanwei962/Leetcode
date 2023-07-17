/*
 * @lc app=leetcode.cn id=191 lang=golang
 * @lcpr version=21909
 *
 * [191] 位1的个数
 */

// @lc code=start
package main

func hammingWeight(n uint32) (ans int) {
	for n > 0 {
		n &= n - 1
		ans++
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 00000000000000000000000000001011\n
// @lcpr case=end

// @lcpr case=start
// 00000000000000000000000010000000\n
// @lcpr case=end

// @lcpr case=start
// 11111111111111111111111111111101\n
// @lcpr case=end

*/
