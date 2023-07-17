/*
 * @lc app=leetcode.cn id=318 lang=golang
 * @lcpr version=21909
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
// 没有相等的-> i&j==0 位运算
package main

func maxProduct(words []string) (ans int) {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// ["abcw","baz","foo","bar","xtfn","abcdef"]\n
// @lcpr case=end

// @lcpr case=start
// ["a","ab","abc","d","cd","bcd","abcd"]\n
// @lcpr case=end

// @lcpr case=start
// ["a","aa","aaa","aaaa"]\n
// @lcpr case=end

*/
