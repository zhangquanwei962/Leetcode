/*
 * @lc app=leetcode.cn id=543 lang=golang
 * @lcpr version=21909
 *
 * [543] 二叉树的直径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 后序
package main

var maxSum int

func diameterOfBinaryTree(root *TreeNode) int {
	maxSum = -1
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := dfs(node.Left)
		r := dfs(node.Right)

		priceSum := l + r + 1
		maxSum = max(maxSum, priceSum)

		return max(l, r) + 1

	}

	dfs(root)

	return maxSum - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
