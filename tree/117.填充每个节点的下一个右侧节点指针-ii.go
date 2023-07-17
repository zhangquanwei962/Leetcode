/*
 * @lc app=leetcode.cn id=117 lang=golang
 * @lcpr version=21909
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
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
package main

func connect(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		do := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for o := start; o != nil; o = o.Next {
			do(o.Left)
			do(o.Right)
		}
		start = nextStart
	}
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,null,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
