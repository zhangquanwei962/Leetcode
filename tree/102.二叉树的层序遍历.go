/*
 * @lc app=leetcode.cn id=102 lang=golang
 * @lcpr version=21909
 *
 * [102] 二叉树的层序遍历
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

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}

	for i := 0; len(q) > 0; i++ {
		tmp := q
		q = nil
		res = append(res, []int{})

		for _, p := range tmp {
			res[i] = append(res[i], p.Val)
			if p.Left != nil {
				q = append(q, p.Left)
			}
			if p.Right != nil {
				q = append(q, p.Right)
			}
		}
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
