/*
 * @lc app=leetcode.cn id=104 lang=golang
 * @lcpr version=21908
 *
 * [104] 二叉树的最大深度
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
package main

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	lDepth := maxDepth(root.Left)
// 	rDepth := maxDepth(root.Right)
// 	return max(lDepth, rDepth) + 1
// }

func maxDepth(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, cnt int) {
		if node == nil {
			ans = max(cnt, ans)
			return
		}

		cnt++
		// ans = max(ans, cnt)
		dfs(node.Left, cnt)
		dfs(node.Right, cnt)
	}
	dfs(root, 0)
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
