/*
 * @lc app=leetcode.cn id=938 lang=golang
 * @lcpr version=21909
 *
 * [938] 二叉搜索树的范围和
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

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high) + root.Val
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,15,3,7,null,18]\n7\n15\n
// @lcpr case=end

// @lcpr case=start
// [10,5,15,3,7,13,18,1,null,6]\n6\n10\n
// @lcpr case=end

*/
