/*
 * @lc app=leetcode.cn id=129 lang=golang
 * @lcpr version=21909
 *
 * [129] 求根节点到叶节点数字之和
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

func dfs(node *TreeNode, preSum int) int {
	if node == nil {
		return 0
	}
	sum := preSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}

	return dfs(node.Left, sum) + dfs(node.Right, sum)
}
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,0,5,1]\n
// @lcpr case=end

*/
