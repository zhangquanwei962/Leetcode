/*
 * @lc app=leetcode.cn id=2716 lang=golang
 * @lcpr version=21909
 *
 * [2716] 最小化字符串长度
 */

// @lc code=start
// O(N) O(1)
package main

import "math/bits"

func minimizedStringLength(s string) int {
	mask := uint(0)
	for _, c := range s {
		mask |= 1 << uint(c-'a')
	}
	return bits.OnesCount(mask)
}

// @lc code=end

/*
// @lcpr case=start
// "aaabc"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

// @lcpr case=start
// "dddaaa"\n
// @lcpr case=end

*/
