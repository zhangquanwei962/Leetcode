/*
 * @lc app=leetcode.cn id=1372 lang=golang
 * @lcpr version=21909
 *
 * [1372] 二叉树中的最长交错路径
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

/*
	if node==nil{
		return 1 // 点数
		return 0 // 路径高度
	}
*/
package main

func longestZigZag(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		_, lr := dfs(node.Left)
		rl, _ := dfs(node.Right)
		ans = max(ans, max(rl, lr))
		return 1 + lr, 1 + rl
	}
	dfs(root)
	return
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
// [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,null,1,null,null,1,1,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
