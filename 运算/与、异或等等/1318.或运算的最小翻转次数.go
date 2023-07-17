/*
 * @lc app=leetcode.cn id=1318 lang=golang
 * @lcpr version=21909
 *
 * [1318] 或运算的最小翻转次数
 */

// @lc code=start
// 10^9 < 2^31-1
// O(logC)
package main

func minFlips(a int, b int, c int) (ans int) {
	for i := 0; i < 31; i++ {
		bit_a := (a >> i) & 1
		bit_b := (b >> i) & 1
		bit_c := (c >> i) & 1
		if bit_c == 0 {
			ans += bit_a + bit_b
		} else if bit_b|bit_a == 0 {
			ans += 1
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 2\n6\n5\n
// @lcpr case=end

// @lcpr case=start
// 4\n2\n7\n
// @lcpr case=end

// @lcpr case=start
// 1\n2\n3\n
// @lcpr case=end

*/
