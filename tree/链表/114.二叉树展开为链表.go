/*
 * @lc app=leetcode.cn id=114 lang=golang
 * @lcpr version=21909
 *
 * [114] 二叉树展开为链表
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

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

// func flatten(root *TreeNode)  {
//     list := []*TreeNode{}
//     stack := []*TreeNode{}
//     node := root
//     for node != nil || len(stack) > 0 {
//         for node != nil {
//             list = append(list, node)
//             stack = append(stack, node)
//             node = node.Left
//         }
//         node = stack[len(stack)-1]
//         node = node.Right
//         stack = stack[:len(stack)-1]
//     }

//     for i := 1; i < len(list); i++ {
//         prev, curr := list[i-1], list[i]
//         prev.Left, prev.Right = nil, curr
//     }
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,5,3,4,null,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
