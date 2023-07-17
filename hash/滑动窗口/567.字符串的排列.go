/*
 * @lc app=leetcode.cn id=567 lang=golang
 * @lcpr version=21909
 *
 * [567] 字符串的排列
 */

// @lc code=start
// 同438
package main

func checkInclusion(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	if sLen > pLen {
		return false
	}

	var sCount, pCount [26]int
	for i, ch := range s {
		sCount[ch-'a']++
		pCount[p[i]-'a']++
	}
	if sCount == pCount {
		return true
	}

	for i, ch := range p[:pLen-sLen] {
		pCount[ch-'a']--
		pCount[p[i+sLen]-'a']++
		if sCount == pCount {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "eidbaooo"\n
// @lcpr case=end

// @lcpr case=start
// "eidboaoo"\n
// @lcpr case=end

*/
