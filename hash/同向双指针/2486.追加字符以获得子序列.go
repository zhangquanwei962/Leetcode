/*
 * @lc app=leetcode.cn id=2486 lang=golang
 * @lcpr version=21909
 *
 * [2486] 追加字符以获得子序列
 */

// @lc code=start
// 两个子序列，双指针匹配，刻不容缓
// string的每一个都是rune，而直接t[j]
// 是byte
package main

func appendCharacters(s string, t string) int {
	j, m := 0, len(t)
	for _, c := range s {
		if byte(c) == t[j] {
			j++
			if j == m {
				return 0
			}
		}
	}
	return m - j
}

// @lc code=end

/*
// @lcpr case=start
// "coaching"\n"coding"\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "z"\n"abcde"\n
// @lcpr case=end

*/
