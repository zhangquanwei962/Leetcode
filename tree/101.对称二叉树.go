/*
 * @lc app=leetcode.cn id=101 lang=golang
 * @lcpr version=21909
 *
 * [101] 对称二叉树
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

func isSymmetric(root *TreeNode) bool {
	var isSameTree func(*TreeNode, *TreeNode) bool
	isSameTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
	}
	return isSameTree(root.Left, root.Right)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,3,4,4,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,null,3,null,3]\n
// @lcpr case=end

*/
