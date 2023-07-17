/*
 * @lc app=leetcode.cn id=998 lang=golang
 * @lcpr version=21907
 *
 * [998] 最大二叉树 II
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
*
或者可以递归，节点值<插入值（或者为空），插入节点作为当前节点的gen
当前节点作为插入结点的左
节点值>插入值，插入值递归插入当前节点的右子树
*/

//	func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
//	    if val > root.Val {
//	        node := &TreeNode{Val: val, Left: root}
//	        return node
//	    }
//	    if root.Right == nil {
//	        root.Right = &TreeNode{Val: val}
//	        return root
//	    }
//	    root.Right = insertIntoMaxTree(root.Right, val)
//	    return root
//	}
package main

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	if val > root.Val {
		node.Left = root
		return node
	}
	cur := root
	for cur.Right != nil && val < cur.Right.Val {
		cur = cur.Right
	}
	node.Left = cur.Right
	cur.Right = node
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [4,1,3,null,null,2]\n5\n
// @lcpr case=end

// @lcpr case=start
// [5,2,4,null,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [5,2,3,null,1]\n4\n
// @lcpr case=end

*/
