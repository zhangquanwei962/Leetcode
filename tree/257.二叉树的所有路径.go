/*
 * @lc app=leetcode.cn id=257 lang=golang
 * @lcpr version=21909
 *
 * [257] 二叉树的所有路径
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
// 传递进去path
package main

import "strconv"

func dfs(node *TreeNode, path string, res *[]string) {
	if node == nil {
		return
	}

	path += strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		*res = append(*res, path)
		return
	}

	dfs(node.Left, path+"->", res)
	dfs(node.Right, path+"->", res)
	return
}
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	dfs(root, "", &res)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
