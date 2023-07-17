/*
 * @lc app=leetcode.cn id=72 lang=golang
 * @lcpr version=21908
 *
 * [72] 编辑距离
 */

// @lc code=start
package main

func minDistance(s string, t string) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = j
	}

	for i, x := range s {
		f[i+1][0] = i + 1
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(min(f[i][j+1], f[i+1][j]), f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func minDistance(s, t string) int {
//     n, m := len(s), len(t)
//     memo := make([][]int, n)
//     for i := range memo {
//         memo[i] = make([]int, m)
//         for j := range memo[i] {
//             memo[i][j] = -1 // -1 表示还没有计算过
//         }
//     }
//     var dfs func(int, int) int
//     dfs = func(i, j int) (res int) {
//         if i < 0 {
//             return j + 1
//         }
//         if j < 0 {
//             return i + 1
//         }
//         p := &memo[i][j]
//         if *p != -1 {
//             return *p
//         }
//         defer func() { *p = res }()
//         if s[i] == t[j] {
//             return dfs(i-1, j-1)
//         }
//         return min(min(dfs(i-1, j), dfs(i, j-1)), dfs(i-1, j-1)) + 1
//     }
//     return dfs(n-1, m-1)
// }

// func min(a, b int) int { if b < a { return b }; return a }

// @lc code=end

/*
// @lcpr case=start
// "horse"\n"ros"\n
// @lcpr case=end

// @lcpr case=start
// "intention"\n"execution"\n
// @lcpr case=end

*/
