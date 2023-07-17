/*
 * @lc app=leetcode.cn id=1371 lang=golang
 * @lcpr version=21909
 *
 * [1371] 每个元音包含偶数次的最长子字符串
 */

// @lc code=start
package main

import (
	"math"
)

var DICT = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func findTheLongestSubstring(s string) int {
	cnt := make(map[int]int)
	cnt[0] = -1
	ans, mask := 0, 0
	for i := 0; i < len(s); i++ {
		if _, ok := DICT[s[i]]; ok {
			mask ^= 1 << (s[i] - 'a')
		}
		if v, ok := cnt[mask]; ok {
			ans = int(math.Max(float64(i-v), float64(ans)))
		} else {
			cnt[mask] = i
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "eleetminicoworoep"\n
// @lcpr case=end

// @lcpr case=start
// "leetcodeisgreat"\n
// @lcpr case=end

// @lcpr case=start
// "bcbcbc"\n
// @lcpr case=end

*/
