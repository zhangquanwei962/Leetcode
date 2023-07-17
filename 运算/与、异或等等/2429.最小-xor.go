/*
 * @lc app=leetcode.cn id=2429 lang=golang
 * @lcpr version=21909
 *
 * [2429] 最小 XOR
 */

// @lc code=start

// 贪心，c2是x 的1的个数
// 能影响异或结果的只有1，异或是不进位加法
// 为了让异或和尽量小，这些1应当从高位到低位匹配num1中的
// 1；如果匹配完了还有多余的
// 1，那么就从低位到高位把 0改成 1 __builtin_popcount
// 因为是从num1反推x
package main

import "math/bits"

func minimizeXor(num1 int, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))

	for ; c2 < c1; c2++ {
		num1 &= num1 - 1 // 最低的1变0
	}

	for ; c2 > c1; c2-- {
		num1 |= num1 + 1 //最低的0变1
	}
	return num1
}

// @lc code=end

/*
// @lcpr case=start
// 3\n5\n
// @lcpr case=end

// @lcpr case=start
// 1\n12\n
// @lcpr case=end

*/
