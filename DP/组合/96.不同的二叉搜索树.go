/*
 * @lc app=leetcode.cn id=96 lang=golang
 * @lcpr version=21909
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
package main

func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 左子树  右子树
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
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
