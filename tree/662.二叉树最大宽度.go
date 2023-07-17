/*
 * @lc app=leetcode.cn id=662 lang=golang
 * @lcpr version=21909
 *
 * [662] 二叉树最大宽度
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
func widthOfBinaryTree(root *TreeNode) int {
	levelMin := map[int]int{}
	var dfs func(*TreeNode, int, int) int
	dfs = func(node *TreeNode, depth, index int) int {
		if node == nil {
			return 0
		}
		if _, ok := levelMin[depth]; !ok {
			levelMin[depth] = index // 每一层最先访问到的节点会是最左边的节点，即每一层编号的最小值
		}
		return max(index-levelMin[depth]+1, max(dfs(node.Left, depth+1, index<<1), dfs(node.Right, depth+1, index<<1+1)))
	}
	return dfs(root, 1, 1)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,2,5,3,null,9]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2,5,null,null,9,6,null,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2,5]\n
// @lcpr case=end

*/

