/*
 * @lc app=leetcode.cn id=988 lang=golang
 * @lcpr version=21909
 *
 * [988] 从叶结点开始的最小字符串
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
// 后序遍历 XXX，只能暴力
package main

import (
	"sort"
)

func smallestFromLeaf(root *TreeNode) string {
	var path []string
	dfs(root, "", &path)
	sort.Strings(path)
	return path[0]
}

func dfs(root *TreeNode, s string, path *[]string) {
	if root == nil {
		return
	}
	s = string('a'+root.Val) + s
	if root.Left == nil && root.Right == nil {
		// rev := []rune(s)
		// for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		// 	rev[i], rev[j] = rev[j], rev[i]
		// }
		*path = append(*path, s)
		return
	}
	dfs(root.Left, s, path)
	dfs(root.Right, s, path)
}

// func dfs(node *TreeNode) string {
// 	if node == nil {
// 		return ""
// 	}

// 	str := string('a' + node.Val)
// 	l := dfs(node.Left)
// 	r := dfs(node.Right)

// 	if l == "" && r == "" {
// 		return str
// 	} else if l == "" {
// 		return r + str
// 	} else if r == "" {
// 		return l + str
// 	}
// 	return min(l+str, r+str)
// }
// func smallestFromLeaf(root *TreeNode) string {
// 	var res string
// 	res = dfs(root)
// 	return res
// }
// func min(a, b string) string {
// 	if strings.Compare(a, b) < 0 {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

/*
// @lcpr case=start
// [0,1,2,3,4,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [25,1,3,1,3,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,null,1,0,null,0]\n
// @lcpr case=end

*/
