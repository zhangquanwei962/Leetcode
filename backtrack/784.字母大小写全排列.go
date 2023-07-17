/*
 * @lc app=leetcode.cn id=784 lang=golang
 * @lcpr version=21907
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
package main

func letterCasePermutation(s string) (ans []string) {
	n := len(s)
	var dfs func(int, []byte)
	dfs = func(i int, b []byte) {
		for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
		}
		if i == n {
			ans = append(ans, string(b))
			return
		}

		dfs(i+1, b)

		b[i] ^= 32
		dfs(i+1, b)
		b[i] ^= 32
	}
	dfs(0, []byte(s))
	return
}

// func letterCasePermutation(s string) (ans []string) {
// 	n := len(s)
// 	var dfs func(int, []byte)
// 	dfs = func(i int, b []byte) {
// 		ans = append(ans, string(b))
// 		for j := i; j < n; j++ {
// 			if s[j] >= '0' && s[j] <= '9' {
// 				continue
// 			}
// 			b[j] ^= 32
// 			dfs(j+1, b)
// 			b[j] ^= 32
// 		}
// 	}
// 	dfs(0, []byte(s))
// 	return
// }

// @lc code=end

/*
// @lcpr case=start
// "a1b2"\n
// @lcpr case=end

// @lcpr case=start
// "3z4"\n
// @lcpr case=end

*/
