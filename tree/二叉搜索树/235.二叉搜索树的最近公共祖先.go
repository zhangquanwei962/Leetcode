/*
 * @lc app=leetcode.cn id=235 lang=golang
 * @lcpr version=21909
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	x := root.Val
	if p.Val < x && q.Val < x {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > x && q.Val > x {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n8\n
// @lcpr case=end

// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n4\n
// @lcpr case=end

*/
