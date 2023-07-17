/*
 * @lc app=leetcode.cn id=98 lang=golang
 * @lcpr version=21909
 *
 * [98] 验证二叉搜索树
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

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return math.MaxInt, math.MinInt
	}
	lMin, lMax := dfs(node.Left)
	rMin, rMax := dfs(node.Right)
	x := node.Val
	if x <= lMax || x >= rMin {
		return math.MinInt, math.MaxInt
	}
	return min(lMin, x), max(rMax, x)
}

func isValidBST(root *TreeNode) bool {
	_, mx := dfs(root)
	return mx != math.MaxInt
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 中序

// func isValidBST(root *TreeNode) bool {
//     pre := math.MinInt
//     var dfs func(*TreeNode) bool
//     dfs = func(node *TreeNode) bool {
//         if node == nil {
//             return true
//         }
//         if !dfs(node.Left) || node.Val <= pre {
//             return false
//         }
//         pre = node.Val
//         return dfs(node.Right)
//     }
//     return dfs(root)
// }

// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,4,null,null,3,6]\n
// @lcpr case=end

*/
