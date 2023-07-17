/*
 * @lc app=leetcode.cn id=1373 lang=golang
 * @lcpr version=21909
 *
 * [1373] 二叉搜索子树的最大键值和
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

func maxSumBST(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int, int)
	dfs = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt, 0
		}
		lMin, lMax, lSum := dfs(node.Left)
		rMin, rMax, rSum := dfs(node.Right)
		x := node.Val
		if x <= lMax || x >= rMin { // 不是二叉树
			return math.MinInt, math.MaxInt, 0
		}

		s := lSum + rSum + x
		ans = max(ans, s)
		return min(lMin, x), max(rMax, x), s
	}
	dfs(root)
	return
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

// @lc code=end

/*
// @lcpr case=start
// [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,null,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [-4,-2,-5]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,8,3,null,6,3]\n
// @lcpr case=end

*/
