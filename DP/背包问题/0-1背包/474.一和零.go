/*
 * @lc app=leetcode.cn id=474 lang=golang
 * @lcpr version=21907
 *
 * [474] 一和零
 */

// @lc code=start
package main

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for _, str := range strs {
		c0 := strings.Count(str, "0")
		c1 := len(str) - c0
		for j := m; j >= c0; j-- {
			for k := n; k >= c1; k-- {
				f[j][k] = max(f[j][k], f[j-c0][k-c1]+1)
			}
		}
	}
	return f[m][n]
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
// ["10", "0001", "111001", "1", "0"]\n5\n3\n
// @lcpr case=end

// @lcpr case=start
// ["10", "0", "1"]\n1\n1\n
// @lcpr case=end

*/
