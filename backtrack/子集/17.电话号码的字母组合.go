/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=21909
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
package main

var mapping = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return
	}

	path := make([]byte, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		for _, c := range mapping[digits[i]-'0'] {
			path[i] = byte(c)
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/
