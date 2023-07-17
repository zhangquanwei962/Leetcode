/*
 * @lc app=leetcode.cn id=700 lang=golang
 * @lcpr version=21909
 *
 * [700] 二叉搜索树中的搜索
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

func searchBST(root *TreeNode, val int) *TreeNode {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(tn *TreeNode) *TreeNode {
		if tn == nil {
			return nil
		}
		if val == tn.Val {
			return tn
		}
		if val > tn.Val {
			return dfs(tn.Right)
		}
		return dfs(tn.Left)
	}
	return dfs(root)
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,1,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [4,2,7,1,3]\n5\n
// @lcpr case=end

*/
