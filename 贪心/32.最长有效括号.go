/*
 * @lc app=leetcode.cn id=32 lang=golang
 * @lcpr version=21909
 *
 * [32] 最长有效括号
 */

// @lc code=start
// 贪心 类似于301题
package main

func longestValidParentheses(s string) int {
	left, right, maxLength, n := 0, 0, 0, len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func longestValidParentheses(s string) int {
// 	n := len(s)
// 	dp := make([]int, n)
// 	ans := 0
// 	for i := 1; i < n; i++ {
// 		if s[i] == ')' {
// 			if s[i-1] == '(' {
// 				if i >= 2 {
// 					dp[i] = dp[i-2] + 2
// 				} else {
// 					dp[i] = 2
// 				}
// 			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
// 				if i-dp[i-1]-2 >= 0 {
// 					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
// 				} else {
// 					dp[i] = dp[i-1] + 2
// 				}
// 			}
// 			ans = max(ans, dp[i])
// 		}
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

/*
// @lcpr case=start
// "(()"\n
// @lcpr case=end

// @lcpr case=start
// ")()())"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

*/
