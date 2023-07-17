/*
 * @lc app=leetcode.cn id=226 lang=golang
 * @lcpr version=21909
 *
 * [226] 翻转二叉树
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
// 都反转一下
package main

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(tn *TreeNode) *TreeNode {
		if tn == nil {
			return tn
		}

		left := dfs(tn.Left)
		right := dfs(tn.Right)
		tn.Left = right
		tn.Right = left
		return tn
	}

	return dfs(root)
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,1,3,6,9]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
