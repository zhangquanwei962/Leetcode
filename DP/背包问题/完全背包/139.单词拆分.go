/*
 * @lc app=leetcode.cn id=139 lang=golang
 * @lcpr version=21909
 *
 * [139] 单词拆分
 */

// @lc code=start
// 完全背包
package main

// func wordBreak(s string, wordDict []string) bool {
// 	n := len(s)
// 	seen := make(map[string]bool, len(wordDict))
// 	for _, word := range wordDict {
// 		seen[word] = true
// 	}

// 	var dfs func(int) bool
// 	dfs = func(i int) bool {
// 		if i == n {
// 			return true
// 		}
// 		for j := i + 1; j <= n; j++ {
// 			if seen[s[i:j]] && dfs(j) {
// 				return true
// 			}
// 		}
// 		return false
// 	}

// 	return dfs(0)
// }

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if wordDictSet[s[j:i]] {
				dp[i] = dp[i] || dp[j]
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n["leet", "code"]\n
// @lcpr case=end

// @lcpr case=start
// "applepenapple"\n["apple", "pen"]\n
// @lcpr case=end

// @lcpr case=start
// "catsandog"\n["cats", "dog", "sand", "and", "cat"]\n
// @lcpr case=end

*/
