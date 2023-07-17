/*
 * @lc app=leetcode.cn id=106 lang=golang
 * @lcpr version=21909
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	for i, v := range inorder {
		if v == rootVal {
			return &TreeNode{
				rootVal,
				buildTree(inorder[:i], postorder[:i]),
				buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
			}
		}
	}
	panic(1)
}

// @lc code=end

/*
// @lcpr case=start
// [9,3,15,20,7]\n[9,15,7,20,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
