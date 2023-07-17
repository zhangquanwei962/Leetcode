/*
 * @lc app=leetcode.cn id=2246 lang=golang
 * @lcpr version=21909
 *
 * [2246] 相邻字符不同的最长路径
 */

// @lc code=start
// https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/solution/by-endlesscheng-92fw/
package main

func longestPath(parent []int, s string) (ans int) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		pa := parent[i]
		g[pa] = append(g[pa], i)
	}

	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		for _, y := range g[x] {
			len := dfs(y) + 1
			if s[y] != s[x] {
				ans = max(ans, maxLen+len)
				maxLen = max(maxLen, len)
			}
		}
		return
	}
	dfs(0)
	return ans + 1
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
// [-1,0,0,1,1,2]\n"abacbe"\n
// @lcpr case=end

// @lcpr case=start
// [-1,0,0,0]\n"aabc"\n
// @lcpr case=end

*/
