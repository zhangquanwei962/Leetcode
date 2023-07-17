/*
 * @lc app=leetcode.cn id=103 lang=golang
 * @lcpr version=21909
 *
 * [103] 二叉树的锯齿形层序遍历
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

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	cur := []*TreeNode{root}
	for even := false; len(cur) > 0; even = !even {
		nxt := []*TreeNode{}
		vals := make([]int, len(cur)) // 大小已知
		for i, node := range cur {
			if even {
				vals[len(cur)-1-i] = node.Val // 倒着添加
			} else {
				vals[i] = node.Val
			}
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
		ans = append(ans, vals)
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
