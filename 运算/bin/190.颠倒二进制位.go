/*
 * @lc app=leetcode.cn id=190 lang=golang
 * @lcpr version=21907
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
package main

func reverseBits(num uint32) (res uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		res |= num & 1 << (31 - i)
		num >>= 1
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 00000010100101000001111010011100\n
// @lcpr case=end

// @lcpr case=start
// 11111111111111111111111111111101\n
// @lcpr case=end

*/
