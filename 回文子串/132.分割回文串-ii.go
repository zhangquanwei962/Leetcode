/*
 * @lc app=leetcode.cn id=132 lang=golang
 * @lcpr version=21908
 *
 * [132] 分割回文串 II
 */

// @lc code=start
package main

import "math"

func minCut(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt32
	}

	f[0] = 0

	for i := 0; i < n*2-1; i++ {
		l, r := i/2, i/2+i%2
		f[l+1] = min(f[l+1], f[l]+1) // 不选s[l]
		for l >= 0 && r < n && s[l] == s[r] {
			f[r+1] = min(f[r+1], f[l]+1)
			l--
			r++
		}
	}

	return f[n] - 1
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
// "aab"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n
// @lcpr case=end

*/
