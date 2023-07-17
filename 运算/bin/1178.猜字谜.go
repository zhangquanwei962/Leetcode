/*
 * @lc app=leetcode.cn id=1178 lang=golang
 * @lcpr version=21907
 *
 * [1178] 猜字谜
 */

// @lc code=start
package main

import (
	"math/bits"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
	const puzzlesLength = 7
	cnt := map[int]int{}
	for _, s := range words {
		mask := 0
		for _, ch := range s {
			mask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(uint(mask)) <= puzzlesLength {
			cnt[mask]++
		}
	}

	ans := make([]int, len(puzzles))
	for i, s := range puzzles {
		first := 1 << (s[0] - 'a')
		mask := 0
		for _, ch := range s[1:] {
			mask |= 1 << (ch - 'a')
		}

		subset := mask
		for {
			ans[i] += cnt[first|subset]
			subset = (subset - 1) & mask
			if subset == mask {
				break
			}
		}
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// ["aaaa","asas","able","ability","actt","actor","access"]\n["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]\n
// @lcpr case=end

*/
