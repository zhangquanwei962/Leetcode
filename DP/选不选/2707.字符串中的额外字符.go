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
	f := make([]int, n+1)

	for i := 0; i < n; i++ {
		f[i+1] = f[i] + 1         // 不选
		for j := 0; j <= i; j++ { // 枚举选哪个
			if has[s[j:i+1]] {
				f[i+1] = min(f[i+1], f[j])
			}
		}
	}
	return f[n]

}

func min(a, b int) int {
	if a > b {
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
