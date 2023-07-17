/*
 * @lc app=leetcode.cn id=655 lang=golang
 * @lcpr version=21909
 *
 * [655] 输出二叉树
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
// 二叉树高度
package main

import "strconv"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func calHight(node *TreeNode) (h int) {
// 	if node.Left != nil {
// 		h = calHight(node.Left) + 1
// 	}
// 	if node.Right != nil {
// 		h = max(h, calHight(node.Right)+1)
// 	}
// 	return
// }

func calHight(node *TreeNode) (h int) {
	if node == nil {
		return 0
	}
	l := calHight(node.Left)
	r := calHight(node.Right)
	return max(l, r) + 1
}

func printTree(root *TreeNode) [][]string {
	height := calHight(root) - 1
	m := height + 1
	n := 1<<m - 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(height-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,null,4]\n
// @lcpr case=end

*/
