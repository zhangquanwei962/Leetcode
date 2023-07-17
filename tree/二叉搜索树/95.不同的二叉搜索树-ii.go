/*
 * @lc app=leetcode.cn id=95 lang=golang
 * @lcpr version=21909
 *
 * [95] 不同的二叉搜索树 II
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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(i, j int) []*TreeNode {
	if j < i {
		return []*TreeNode{nil}
	}

	allTreeNode := []*TreeNode{}
	for k := i; k <= j; k++ {
		leftTree := helper(i, k-1)
		rightTree := helper(k+1, j)
		for _, left := range leftTree {
			for _, right := range rightTree {
				curNode := &TreeNode{k, left, right}
				allTreeNode = append(allTreeNode, curNode)
			}
		}
	}
	return allTreeNode
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
