/*
 * @lc app=leetcode.cn id=2673 lang=golang
 * @lcpr version=21909
 *
 * [2673] 使二叉树所有路径值相等的最小代价
 */

// @lc code=start
// 由于这两条路径除了叶子节点不一样，其余节点都一样，所以为了让这两条路径的路径和相等，必须修改叶子节点的值。
package main

func minIncrements(n int, cost []int) (ans int) {
	for i := n / 2; i > 0; i-- {
		left, right := cost[i*2-1], cost[i*2]
		if left > right {
			left, right = right, left
		}
		ans += right - left
		cost[i-1] += right // 累加路径和
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[1,5,2,2,3,3,1]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[5,3,3]\n
// @lcpr case=end

*/
