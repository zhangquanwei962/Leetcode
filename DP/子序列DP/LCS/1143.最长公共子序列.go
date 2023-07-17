/*
 * @lc app=leetcode.cn id=1143 lang=golang
 * @lcpr version=21907
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
package main

func longestCommonSubsequence(s, t string) int {
	// n, m := len(s), len(t)
	// cache := make([][]int, n)
	// for i := range cache {
	// 	cache[i] = make([]int, m)
	// 	for j := range cache[i] {
	// 		cache[i][j] = -1 // -1 表示没有访问过
	// 	}
	// }
	// var dfs func(int, int) int //dfs(i,j)表示前i，j的公共
	// dfs = func(i, j int) (res int) {
	// 	if i < 0 || j < 0 {
	// 		return
	// 	}
	// 	C := &cache[i][j]
	// 	if *C != -1 {
	// 		return *C
	// 	}
	// 	defer func() { *C = res }()
	// 	if s[i] == t[j] {
	// 		return dfs(i-1, j-1) + 1
	// 	}
	// 	return max(dfs(i-1, j), dfs(i, j-1))
	// }
	// return dfs(n-1, m-1)

	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}

	for i, x := range s {
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// func longestCommonSubsequence(s, t string) int {
// 	n, m := len(s), len(t)
// 	f := [2][]int{make([]int, m+1), make([]int, m+1)}
// 	for i, x := range s {
// 		for j, y := range t {
// 			if x == y {
// 				f[(i+1)%2][j+1] = f[i%2][j] + 1
// 			} else {
// 				f[(i+1)%2][j+1] = max(f[i%2][j+1], f[(i+1)%2][j])
// 			}
// 		}
// 	}
// 	return f[n%2][m]
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

/*
// @lcpr case=start
// "abcde"\n"ace"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"def"\n
// @lcpr case=end

*/
