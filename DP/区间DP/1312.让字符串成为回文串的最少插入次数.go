/*
 * @lc app=leetcode.cn id=1312 lang=golang
 * @lcpr version=21909
 *
 * [1312] 让字符串成为回文串的最少插入次数
 */

// @lc code=start
package main

func minInsertions(s string) int {
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

	return n - dfs(0, n-1)
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
// "zzazz"\n
// @lcpr case=end

// @lcpr case=start
// "mbadm"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n
// @lcpr case=end

*/
