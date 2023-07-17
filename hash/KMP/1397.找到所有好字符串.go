/*
 * @lc app=leetcode.cn id=1397 lang=golang
 * @lcpr version=21909
 *
 * [1397] 找到所有好字符串
 */

// @lc code=start
package main

const MOD = 1e9 + 7

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	m := len(evil)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// for i, j := 1, -1; i < m; i++ {
	// 	for j > -1 && evil[i] != evil[j+1] {
	// 		j = ne[j]
	// 	}
	// 	if evil[i] == evil[j+1] {
	// 		j++
	// 	}
	// 	ne[i] = j
	// }
	ne := make([]int, m+1)
	ne[0] = -1
	l, r := -1, 0

	for r < m {
		for l == -1 || (l < m-1 && r < m && evil[l] == evil[r]) {
			l++
			r++
			ne[r] = l
		}

		l = ne[l]
	}

	var f func(int, int, bool, bool) int
	f = func(i, ei int, equal_1, equal_2 bool) (res int) {
		if ei == m {
			return 0
		}
		if i == n {
			return 1
		}

		if !equal_1 && !equal_2 {
			dv := &dp[i][ei]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}

		down := byte('a')
		if equal_1 {
			down = s1[i]
		}
		up := byte('z')
		if equal_2 {
			up = s2[i]
		}

		for d := down; d <= up; d++ {
			j := ei
			// KMP匹配，将 ei 调整到合适位置
			for j != -1 && d != evil[j] {
				j = ne[j]
			}
			res += f(i+1, j+1, equal_1 && d == s1[i], equal_2 && d == s2[i])
			res %= MOD
		}
		return
	}
	return f(0, 0, true, true)
}

// @lc code=end

/*
// @lcpr case=start
// 2\n"aa"\n"da"\n"b"\n
// @lcpr case=end

// @lcpr case=start
// 8\n"leetcode"\n"leetgoes"\n"leet"\n
// @lcpr case=end

// @lcpr case=start
// 2\n"gx"\n"gz"\n"x"\n
// @lcpr case=end

*/
