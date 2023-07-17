/*
 * @lc app=leetcode.cn id=1261 lang=golang
 * @lcpr version=21909
 *
 * [1261] 在受污染的二叉树中查找元素
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 对于从0开始的“完全二叉树”，符合bit的最低位舍去，其他的
//作为行走路线 0->left 1->right
package main

import (
	"strconv"
)

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	return FindElements{root: root}
}

func (this *FindElements) Find(target int) bool {
	if target < 0 {
		return false
	}
	target++

	bits, cur := strconv.FormatInt(int64(target), 2), this.root
	for i := 1; i < len(bits); i++ {
		if bits[i] == '0' {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
		if cur == nil {
			return false
		}
	}
	return true
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
// @lc code=end

/*
// @lcpr case=start
// ["FindElements","find","find"][[[-1,null,-1]],[1],[2]]\n
// @lcpr case=end

// @lcpr case=start
// ["FindElements","find","find","find"][[[-1,-1,-1,-1,-1]],[1],[3],[5]]\n
// @lcpr case=end

// @lcpr case=start
// ["FindElements","find","find","find","find"][[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]\n
// @lcpr case=end

*/
