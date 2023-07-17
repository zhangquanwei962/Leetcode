/*
 * @lc app=leetcode.cn id=2718 lang=golang
 * @lcpr version=21909
 *
 * [2718] 查询后矩阵的和
 */

// @lc code=start
// 思考 离线思想
// 修改多次只有最后一次修改起作用！！！
// 正难则反
package main

func matrixSumQueries(n int, queries [][]int) (ans int64) {
	vis := [2]map[int]bool{{}, {}}
	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		tp, index, val := q[0], q[1], q[2]
		if !vis[tp][index] { // 后面（>i）没有对这一行/列的操作
			// 这一行/列还剩下 n-len(vis[tp^1]) 个可以填入的格子
			ans += int64(n-len(vis[tp^1])) * int64(val)
			vis[tp][index] = true // 标记操作过
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 3\n[[0,0,1],[1,2,2],[0,2,3],[1,0,4]]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[[0,0,4],[0,1,2],[1,0,1],[0,2,3],[1,2,1]]\n
// @lcpr case=end

*/
