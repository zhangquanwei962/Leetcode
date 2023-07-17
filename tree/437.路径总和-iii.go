/*
 * @lc app=leetcode.cn id=437 lang=golang
 * @lcpr version=21909
 *
 * [437] 路径总和 III
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

var count int

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func dfs(node *TreeNode, sum int) (res int) {
	if node == nil {
		return
	}
	sum -= node.Val
	if sum == 0 {
		res++
	}
	res += dfs(node.Left, sum)
	res += dfs(node.Right, sum)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,-3,3,2,null,11,3,-2,null,1]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

*/
