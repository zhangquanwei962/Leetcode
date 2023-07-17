/*
 * @lc app=leetcode.cn id=337 lang=golang
 * @lcpr version=21909
 *
 * [337] 打家劫舍 III
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

func dfs(node *TreeNode) (int, int) {
	if node == nil { // 边界条件
		return 0, 0 // 没有节点，怎么选都是 0
	}
	lRob, lNotRob := dfs(node.Left)                   // 递归左子树
	rRob, rNotRob := dfs(node.Right)                  // 递归右子树
	rob := lNotRob + rNotRob + node.Val               // 选
	notRob := max(lRob, lNotRob) + max(rRob, rNotRob) // 不选
	return rob, notRob
}

func rob(root *TreeNode) int {
	return max(dfs(root)) // 根节点选或不选的最大值
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
// [3,2,3,null,3,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [3,4,5,1,3,null,1]\n
// @lcpr case=end

*/
