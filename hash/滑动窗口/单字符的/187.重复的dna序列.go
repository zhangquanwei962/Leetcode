/*
 * @lc app=leetcode.cn id=187 lang=golang
 * @lcpr version=21908
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
package main

func findRepeatedDnaSequences(S string) (ans []string) {
	left := 0
	cnt := map[string]int{}
	var s string
	for right, ch := range S {
		s += string(ch)
		if right-left+1 > 10 {
			s = s[1:]
			left++
		}
		if right-left+1 == 10 {
			cnt[s]++
			if cnt[s] == 2 { //避免重复
				ans = append(ans, s)
			}

		}

	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"\n
// @lcpr case=end

// @lcpr case=start
// "AAAAAAAAAAAAA"\n
// @lcpr case=end

*/
