/*
 * @lc app=leetcode.cn id=1163 lang=golang
 * @lcpr version=21909
 *
 * [1163] 按字典序排在最后的子串
 */

// @lc code=start
// 字典序，有单调性，考虑双指针
// 肯定是后缀串 O(n) O(1)
package main

func lastSubstring(s string) string {
	i, n := 0, len(s)
	for j, k := 1, 0; j+k < n; {
		if s[k+j] == s[i+k] {
			k++
		} else if s[i+k] < s[j+k] {
			i += k + 1
			k = 0
			if i >= j {
				j = i + 1
			}
		} else {
			j += k + 1 // 它们的字典序一定小于对应的
			k = 0
		}
	}
	return s[i:]
}

// @lc code=end

/*
// @lcpr case=start
// "abab"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n
// @lcpr case=end

*/
