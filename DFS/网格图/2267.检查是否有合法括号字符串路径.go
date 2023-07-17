/*
 * @lc app=leetcode.cn id=2267 lang=golang
 * @lcpr version=21909
 *
 * [2267] 检查是否有合法括号字符串路径
 */

// @lc code=start
// dfs确定状态+疯狂剪枝
// 括号就是c记录个数，合法就是c任意时刻c>=0并且最后c==0
//从起点到终点，往下走的次数是固定的，即m−1 次，往右走的次数也是固定的，即

// n−1 次，因此路径长度（字符串长度）是一个定值，即

//(m−1)+(n−1)+1=m+n−1。

package main

func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n)&1 == 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' { // 剪枝
		return false
	}

	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, (m+n+1)/2)
		}
	}
	var dfs func(x, y, c int) bool
	dfs = func(x, y, c int) bool {
		if c >= m-x+n-y { // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
			return false
		}
		if x == m-1 && y == n-1 { // 终点
			return c == 1 // 终点一定是 ')'
		}
		if vis[x][y][c] { // 重复访问
			return false
		}
		vis[x][y][c] = true
		if grid[x][y] == '(' {
			c++
		} else if c--; c < 0 { // 非法括号字符串
			return false
		}
		return x < m-1 && dfs(x+1, y, c) || y < n-1 && dfs(x, y+1, c) // 往下或者往右
	}
	return dfs(0, 0, 0) // 起点
}

// @lc code=end

/*
// @lcpr case=start
// [["(","(","("],[")","(",")"],["(","(",")"],["(","(",")"]]\n
// @lcpr case=end

// @lcpr case=start
// [[")",")"],["(","("]]\n
// @lcpr case=end

*/
