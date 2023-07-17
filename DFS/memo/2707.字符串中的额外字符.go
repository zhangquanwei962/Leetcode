/*
 * @lc app=leetcode.cn id=2707 lang=golang
 * @lcpr version=21909
 *
 * [2707] 字符串中的额外字符
 */

// @lc code=start
package main

func minExtraChar(s string, dictionary []string) int {
	has := map[string]bool{}
	for _, s := range dictionary {
		has[s] = true
	}

	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}

		// 不选
		res := dfs(i-1) + 1
		// 枚举选哪个
		for j := 0; j <= i; j++ {
			if has[s[j:i+1]] {
				res = min(res, dfs(j-1))
			}
		}
		*p = res
		return res
	}
	return dfs(n - 1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "leetscode"\n["leet","code","leetcode"]\n
// @lcpr case=end

// @lcpr case=start
// "sayhelloworld"\n["hello","world"]\n
// @lcpr case=end

*/
