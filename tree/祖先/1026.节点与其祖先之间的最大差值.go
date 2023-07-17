/*
 * @lc app=leetcode.cn id=1026 lang=golang
 * @lcpr version=21907
 *
 * [1026] 节点与其祖先之间的最大差值
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// func maxAncestorDiff(root *TreeNode) (ans int) {
// 	var dfs func(*TreeNode, int, int)

// 	dfs = func(node *TreeNode, mn, mx int) {
// 		if node == nil {
// 			return
// 		}

// 		mn = min(mn, node.Val)
// 		mx = max(mx, node.Val)
// 		ans = max(ans, max(node.Val-mn, mx-node.Val))
// 		dfs(node.Left, mn, mx)
// 		dfs(node.Right, mn, mx)
// 	}
// 	dfs(root, root.Val, root.Val)
// 	return
// }

func maxAncestorDiff(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int, int)

	dfs = func(tn *TreeNode, i1, i2 int) {
		if tn == nil {
			ans = max(ans, i2-i1)
			return
		}

		i1 = min(i1, tn.Val)
		i2 = max(i2, tn.Val)
		dfs(tn.Left, i1, i2)
		dfs(tn.Right, i1, i2)
	}
	dfs(root, root.Val, root.Val)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [8,3,10,1,6,null,14,null,null,4,7,13]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2,null,0,3]\n
// @lcpr case=end

*/
