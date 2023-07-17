/*
 * @lc app=leetcode.cn id=111 lang=golang
 * @lcpr version=21909
 *
 * [111] 二叉树的最小深度
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 {
		return r + 1
	} else if r == 0 {
		return l + 1
	}
	return min(l, r) + 1
}

// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
// 	minD := math.MaxInt32

// 	if root.Left != nil {
// 		minD = min(minD, minDepth(root.Left))
// 	}

// 	if root.Right != nil {
// 		minD = min(minD, minDepth(root.Right))
// 	}
// 	return minD + 1
// }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,null,3,null,4,null,5,null,6]\n
// @lcpr case=end

*/
