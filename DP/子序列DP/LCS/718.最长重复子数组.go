/*
 * @lc app=leetcode.cn id=718 lang=golang
 * @lcpr version=21909
 *
 * [718] 最长重复子数组
 */

// @lc code=start
// O(mn) O(mn)
package main

func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	// for i := n - 1; i >= 0; i-- {
	// 	for j := m - 1; j >= 0; j-- {
	// 		if A[i] == B[j] {
	// 			dp[i][j] = dp[i+1][j+1] + 1
	// 		} else {
	// 			dp[i][j] = 0
	// 		}
	// 		if ans < dp[i][j] {
	// 			ans = dp[i][j]
	// 		}
	// 	}
	// }
	for i, v := range A {
		for j, w := range B {
			if v == w {
				dp[i+1][j+1] = dp[i][j] + 1
			}
			if ans < dp[i+1][j+1] {
				ans = dp[i+1][j+1]
			}
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,2,1]\n[3,2,1,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,0,0]\n[0,0,0,0,0]\n
// @lcpr case=end

*/
