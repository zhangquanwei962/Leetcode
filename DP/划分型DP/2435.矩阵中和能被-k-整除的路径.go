/*
 * @lc app=leetcode.cn id=2435 lang=golang
 * @lcpr version=21907
 *
 * [2435] 矩阵中和能被 K 整除的路径
 */

// @lc code=start
/*
 * 主要思路就是把路径和摸k的结果当成一个扩展维度
 */
package main

func numberOfPaths(grid [][]int, k int) int {
	const mod int = 1e9 + 7
	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k)
		}
	}

	f[0][1][0] = 1
	for i, row := range grid {
		for j, x := range row {
			for v := 0; v < k; v++ {
				f[i+1][j+1][(v+x)%k] = (f[i+1][j][v] + f[i][j+1][v]) % mod
			}
		}
	}
	return f[m][n][0]
}

// @lc code=end

/*
// @lcpr case=start
// [[5,2,4],[3,0,5],[0,7,2]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[0,0]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[7,3,4,9],[2,3,6,2],[2,3,7,0]]\n1\n
// @lcpr case=end

*/
