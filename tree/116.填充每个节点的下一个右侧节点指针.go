/*
 * @lc app=leetcode.cn id=116 lang=golang
 * @lcpr version=21909
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// 前序遍历
package main

func connect(root *Node) *Node {
	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}

		// if node.Left != nil && node.Right != nil {
		if node.Left != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			} else {
				node.Right.Next = nil
			}
			dfs(node.Left)
			dfs(node.Right)
		}
		return node
	}
	return dfs(root)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
