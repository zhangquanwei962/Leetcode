/*
 * @lc app=leetcode.cn id=2430 lang=golang
 * @lcpr version=21909
 *
 * [2430] 对字母串可执行的最大删除数
 */

// @lc code=start
// O(nlogn) O(n)
package main

const P uint64 = 131

func deleteString(s string) int {
	n := len(s)
	if allEqual(s) {
		return n // 如果所有字符都相同，直接返回字符串长度
	}

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

	f := make([]int, n+1) // f[i]表示删除str[i...n]所需的最大操作数

	for i := n; i > 0; i-- {
		f[i] = 1 // 删除整个字符串的操作数
		for length := 1; i+2*length-1 <= n; length++ {
			j := i + length - 1
			if get(i, j) == get(j+1, j+length) {
				f[i] = max(f[i], f[j+1]+1)
			}
		}
	}

	return f[1]

}

// func deleteString(s string) int {
// 	n := len(s)
// 	if allEqual(s) { // 特判全部相同的情况
// 		return n
// 	}
// 	lcp := make([][]int, n+1) // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
// 	lcp[n] = make([]int, n+1)
// 	for i := n - 1; i >= 0; i-- {
// 		lcp[i] = make([]int, n+1)
// 		for j := n - 1; j > i; j-- {
// 			if s[i] == s[j] {
// 				lcp[i][j] = lcp[i+1][j+1] + 1
// 			}
// 		}
// 	}
// 	f := make([]int, n)
// 	for i := n - 1; i >= 0; i-- {
// 		for j := 1; i+j*2 <= n; j++ {
// 			if lcp[i][i+j] >= j { // 说明 s[i:i+j] == s[i+j:i+j*2]
// 				f[i] = max(f[i], f[i+j])
// 			}
// 		}
// 		f[i]++
// 	}
// 	return f[0]
// }

func allEqual(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcdabc"\n
// @lcpr case=end

// @lcpr case=start
// "aaabaab"\n
// @lcpr case=end

// @lcpr case=start
// "aaaaa"\n
// @lcpr case=end

*/
