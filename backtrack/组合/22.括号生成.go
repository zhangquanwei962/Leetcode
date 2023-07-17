/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=21909
 *
 * [22] 括号生成
 */

// @lc code=start
package main

func generateParenthesis(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 可以填左括号
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open { // 可以填右括号
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
