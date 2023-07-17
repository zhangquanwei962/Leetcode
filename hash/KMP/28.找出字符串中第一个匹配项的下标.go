/*
 * @lc app=leetcode.cn id=28 lang=golang
 * @lcpr version=21909
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
// O(m+n)  O(m)
package main

func strStr(s string, p string) int {
	n, m := len(s), len(p)
	ne := make([]int, m)
	ne[0] = -1

	for i, j := 1, -1; i < m; i++ {
		for j > -1 && p[i] != p[j+1] {
			j = ne[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		ne[i] = j
	}

	for i, j := 0, -1; i < n; i++ {
		for j > -1 && s[i] != p[j+1] {
			j = ne[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == m-1 {
			// j = ne[m-1]
			return i - m + 1
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// "sadbutsad"\n"sad"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"leeto"\n
// @lcpr case=end

*/
