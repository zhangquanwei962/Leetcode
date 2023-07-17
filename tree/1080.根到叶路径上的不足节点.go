/*
 * @lc app=leetcode.cn id=1080 lang=golang
 * @lcpr version=21908
 *
 * [1080] 根到叶路径上的不足节点
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
// 后序遍历
package main

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}

	// 如果当前节点是叶子节点，则返回 nil 代表该节点需要被删除
	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			return nil
		} else {
			return root
		}
	}

	// 对左右子树进行递归操作
	root.Left = sufficientSubset(root.Left, limit-root.Val)
	root.Right = sufficientSubset(root.Right, limit-root.Val)

	// 如果左右子树都需要被删除，则返回 nil 代表该节点需要被删除
	if root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14]\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,4,8,11,null,17,4,7,1,null,null,5,3]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,-3,-5,null,4,null]\n-1\n
// @lcpr case=end

*/
