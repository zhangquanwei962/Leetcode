/*
 * @lc app=leetcode.cn id=796 lang=golang
 * @lcpr version=21909
 *
 * [796] 旋转字符串
 */

// @lc code=start
package main

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	p := goal[:]
	s = s + s
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

	for i, j := 1, -1; i < n; i++ {
		for j > -1 && s[i] != p[j+1] {
			j = ne[j]
		}
		if s[i] == p[j+1] {
			j++

		}
		if j == m-1 {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abcde"\n"cdeab"\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n"abced"\n
// @lcpr case=end

*/
