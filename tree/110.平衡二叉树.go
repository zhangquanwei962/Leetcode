/*
 * @lc app=leetcode.cn id=110 lang=golang
 * @lcpr version=21909
 *
 * [110] 平衡二叉树
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

func isBalanced(root *TreeNode) bool {
	var getHeight func(*TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := getHeight(node.Left)
		if leftHeight == -1 {
			return -1 // 提前退出，不再递归
		}

		rightHeight := getHeight(node.Right)
		if rightHeight == -1 || abs(rightHeight-leftHeight) > 1 {
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return getHeight(root) != -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,3,3,null,null,4,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
