/*
 * @lc app=leetcode.cn id=1392 lang=golang
 * @lcpr version=21909
 *
 * [1392] 最长快乐前缀
 */

// @lc code=start
// O(N) O(N)
// h
package main

const P uint64 = 131

func longestPrefix(s string) string {
	n := len(s)
	h := make([]uint64, n+1)
	p := make([]uint64, n+1)
	p[0] = 1
	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*P + uint64(s[i-1])
		p[i] = p[i-1] * P
	}

	get := func(l, r int) uint64 {
		return h[r] - h[l-1]*p[r-l+1]
	}

	mx := 0
	for i := 1; i < n; i++ {
		// if h[i] == h[n]-h[n-i]*p[i] {
		if get(1, i) == get(n-i+1, n) {
			mx = i
		}
	}
	return s[:mx]
}

// func longestPrefix(s string) string {
// 	n := len(s)
// 	ne := make([]int, n)
// 	ne[0] = -1
// 	for i, j := 1, -1; i < n; i++ {
// 		for j > -1 && s[i] != s[j+1] {
// 			j = ne[j]
// 		}
// 		if s[i] == s[j+1] {
// 			j++
// 		}
// 		ne[i] = j
// 	}
// 	return s[0 : ne[n-1]+1]
// }

// @lc code=end

/*
// @lcpr case=start
// "level"\n
// @lcpr case=end

// @lcpr case=start
// "ababab"\n
// @lcpr case=end

*/
