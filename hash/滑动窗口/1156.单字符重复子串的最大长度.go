/*
 * @lc app=leetcode.cn id=1156 lang=golang
 * @lcpr version=21909
 *
 * [1156] 单字符重复子串的最大长度
 */

// @lc code=start
package main

import (
	"math"
)

func maxRepOpt1(text string) int {
	cnt := make(map[rune]int)
	tnt := make(map[byte]int)
	ans := math.MinInt
	left := 0
	for _, x := range text {
		cnt[x]++
	}
	for right, x := range text {
		tnt[byte(x)]++
		for len(tnt) > 2 || (minVal(tnt) > 1 && len(tnt) == 2) {
			tnt[text[left]]--
			if tnt[text[left]] == 0 {
				delete(tnt, text[left])
			}
			left++
		}
		mx := maxKeyVal(tnt)
		ans = max(min(cnt[rune(mx)], right-left+1), ans)
	}
	return ans
}

func minVal(m map[byte]int) int {
	min := math.MaxInt
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func maxKeyVal(m map[byte]int) byte {
	max := math.MinInt
	var maxKey byte
	for k, v := range m {
		if v > max {
			max = v
			maxKey = k
		}
	}
	return maxKey
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// "ababa"\n
// @lcpr case=end

// @lcpr case=start
// "aaabaaa"\n
// @lcpr case=end

// @lcpr case=start
// "aaabbaaa"\n
// @lcpr case=end

// @lcpr case=start
// "aaaaa"\n
// @lcpr case=end

// @lcpr case=start
// "abcdef"\n
// @lcpr case=end

*/
