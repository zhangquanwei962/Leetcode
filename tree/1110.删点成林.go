/*
 * @lc app=leetcode.cn id=1110 lang=golang
 * @lcpr version=21908
 *
 * [1110] 删点成林
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

// 后序遍历
package main

func delNodes(root *TreeNode, to_delete []int) (ans []*TreeNode) {
	set := make(map[int]struct{}, len(to_delete))

	for _, x := range to_delete {
		set[x] = struct{}{}
	}

	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if _, ok := set[node.Val]; !ok {
			return node
		}

		if node.Left != nil {
			ans = append(ans, node.Left)
		}

		if node.Right != nil {
			ans = append(ans, node.Right)
		}

		return nil

	}

	if dfs(root) != nil {
		ans = append(ans, root)
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n[3,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,4,null,3]\n[3]\n
// @lcpr case=end

*/
