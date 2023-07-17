/*
 * @lc app=leetcode.cn id=1638 lang=golang
 * @lcpr version=21907
 *
 * [1638] 统计只差一个字符的子串数目
 */

// 2444 795
// @lc code=start
package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func countSubstrings(s string, t string) (ans int) {
	n, m := len(s), len(t)
	for d := 1 - m; d < n; d++ { // d=i-j, j=i-d
		i := max(d, 0)
		for k0, k1 := i-1, i-1; i < n && i-d < m; i++ {
			if s[i] != t[i-d] {
				k0 = k1 // 上上一个不同
				k1 = i  // 上一个不同
			}
			ans += k1 - k0
		}

	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "aba"\n"baba"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n"bb"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "abe"\n"bbc"\n
// @lcpr case=end

*/
