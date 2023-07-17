/*
 * @lc app=leetcode.cn id=516 lang=golang
 * @lcpr version=21909
 *
 * [516] 最长回文子序列
 */

// @lc code=start
package main

// func longestPalindromeSubseq(s string) int {
//     n := len(s)
//     f := make([][]int, n)
//     for i := range f {
//         f[i] = make([]int, n)
//     }
//     for i := n - 1; i >= 0; i-- {
//         f[i][i] = 1
//         for j := i + 1; j < n; j++ {
//             if s[i] == s[j] {
//                 f[i][j] = f[i+1][j-1] + 2
//             } else {
//                 f[i][j] = max(f[i+1][j], f[i][j-1])
//             }
//         }
//     }
//     return f[0][n-1]
// }

// func max(a, b int) int { if a < b { return b }; return a }

func longestPalindromeSubseq(s string) int {
	n := len(s)
	memo := make([][]int, n)

	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		if s[i] == s[j] {
			return dfs(i+1, j-1) + 2
		}
		return max(dfs(i+1, j), dfs(i, j-1))
	}
	return dfs(0, n-1)
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
// "bbbab"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/
