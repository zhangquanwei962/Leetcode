/*
 * @lc app=leetcode.cn id=1297 lang=golang
 * @lcpr version=21909
 *
 * [1297] 子串的最大出现次数
 */

// @lc code=start
// O(NS) O(NS)
package main

func maxFreq(s string, maxLetters int, minSize int, maxSize int) (ans int) {
	n := len(s)
	occ := make(map[string]int)
	for i := 0; i < n-minSize+1; i++ {
		cur := s[i : i+minSize]
		exist := make(map[rune]bool)
		for _, ch := range cur {
			exist[ch] = true
		}
		if len(exist) <= maxLetters {
			occ[cur]++
			ans = max(ans, occ[cur])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// "aababcaab"\n2\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n1\n3\n3\n
// @lcpr case=end

// @lcpr case=start
// "aabcabcab"\n2\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n2\n3\n3\n
// @lcpr case=end

*/
