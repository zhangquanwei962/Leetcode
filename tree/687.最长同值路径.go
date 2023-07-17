/*
 * @lc app=leetcode.cn id=687 lang=golang
 * @lcpr version=21909
 *
 * [687] 最长同值路径
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

func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}

		L := dfs(tn.Left)
		R := dfs(tn.Right)

		if tn.Left != nil && tn.Val == tn.Left.Val {
			L++
		} else {
			L = 0
		}

		if tn.Right != nil && tn.Val == tn.Right.Val {
			R++
		} else {
			R = 0
		}
		ans = max(ans, L+R)
		return max(L, R)
	}

	if root == nil {
		return 0
	}
	dfs(root)
	return ans
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
// [5,4,5,1,1,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,4,5,4,4,5]\n
// @lcpr case=end

*/
