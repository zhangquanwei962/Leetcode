/*
 * @lc app=leetcode.cn id=140 lang=golang
 * @lcpr version=21909
 *
 * [140] 单词拆分 II
 */

// @lc code=start
package main

import "strings"

func wordBreak(s string, wordDict []string) (ans []string) {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}
	path := []string{}

	var dfs func(int)
	dfs = func(i int) {
		if i == len(s) {
			ans = append(ans, strings.Join(path, " "))
			return
		}

		for j := i; j < len(s); j++ {
			str := s[i : j+1]
			if _, ok := set[str]; ok {
				path = append(path, str)
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// "catsanddog"\n["cat","cats","and","sand","dog"]\n
// @lcpr case=end

// @lcpr case=start
// "pineapplepenapple"\n["apple","pen","applepen","pine","pineapple"]\n
// @lcpr case=end

// @lcpr case=start
// "catsandog"\n["cats","dog","sand","and","cat"]\n
// @lcpr case=end

*/
