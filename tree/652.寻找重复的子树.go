/*
 * @lc app=leetcode.cn id=652 lang=golang
 * @lcpr version=21909
 *
 * [652] 寻找重复的子树
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

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := map[*TreeNode]struct{}{}
	seen := map[string]*TreeNode{}
	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		serial := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if n, ok := seen[serial]; ok {
			repeat[n] = struct{}{}
		} else {
			seen[serial] = node
		}
		return serial
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,null,2,4,null,null,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,3,null,3,null]\n
// @lcpr case=end

*/
