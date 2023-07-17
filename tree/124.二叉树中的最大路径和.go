/*
 * @lc app=leetcode.cn id=124 lang=golang
 * @lcpr version=21909
 *
 * [124] 二叉树中的最大路径和
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

import "math"

func maxPathSum(root *TreeNode) int {
	var maxn int = math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		l := max(dfs(tn.Left), 0)
		r := max(dfs(tn.Right), 0)
		pathWeight := l + r + tn.Val
		maxn = max(pathWeight, maxn)
		return max(l, r) + tn.Val
	}
	dfs(root)
	return maxn
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-10,9,20,null,null,15,7]\n
// @lcpr case=end

*/
