/*
 * @lc app=leetcode.cn id=979 lang=golang
 * @lcpr version=21909
 *
 * [979] 在二叉树中分配硬币
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
// 后序遍历
package main

func distributeCoins(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		d := dfs(node.Left) + dfs(node.Right) + node.Val - 1
		ans += abs(d)
		return d
	}
	dfs(root)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [3,0,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,0]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,0,null,3]\n
// @lcpr case=end

*/
