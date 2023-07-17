/*
 * @lc app=leetcode.cn id=2385 lang=golang
 * @lcpr version=21908
 *
 * [2385] 感染二叉树需要的总时间
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

// 首先这个一看就是个bfs，然后再看这个需要找到每个节点的
// 父节点以及start跑一遍bfs
package main

func amountOfTime(root *TreeNode, start int) int {
	var st *TreeNode
	parents := map[*TreeNode]*TreeNode{}
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node, pa *TreeNode) {
		if node == nil {
			return
		}
		if node.Val == start {
			st = node
		}
		parents[node] = pa
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	ans := -1
	vis := map[*TreeNode]bool{nil: true, st: true} //兼顾存在与为空
	for q := []*TreeNode{st}; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, node := range tmp {
			if node != nil {
				if !vis[node.Left] {
					vis[node.Left] = true
					q = append(q, node.Left)
				}
				if !vis[node.Right] {
					vis[node.Right] = true
					q = append(q, node.Right)
				}
				if p := parents[node]; !vis[p] {
					vis[p] = true
					q = append(q, p)
				}
			}
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,5,3,null,4,10,6,9,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/
