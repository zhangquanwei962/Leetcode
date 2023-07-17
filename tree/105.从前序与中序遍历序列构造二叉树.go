/*
 * @lc app=leetcode.cn id=105 lang=golang
 * @lcpr version=21909
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	for i, v := range inorder {
		if v == rootVal {
			return &TreeNode{
				rootVal,
				buildTree(preorder[1:i+1], inorder[:i]),
				buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	panic("invalid input")
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,15,7]\n[9,3,15,20,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
