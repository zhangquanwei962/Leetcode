/*
 * @lc app=leetcode.cn id=2713 lang=golang
 * @lcpr version=21909
 *
 * [2713] 矩阵中严格递增的单元格数
 */

// @lc code=start
package main

import "sort"

func maxIncreasingCells(mat [][]int) (ans int) {
	type pair struct{ x, y int }
	g := map[int][]pair{}
	for i, row := range mat {
		for j, x := range row {
			g[x] = append(g[x], pair{i, j}) // 相同元素放在同一组，统计位置
		}
	}
	a := make([]int, 0, len(g))
	for k := range g {
		a = append(a, k)
	}
	sort.Ints(a) // 从小到大

	rowMax := make([]int, len(mat))
	colMax := make([]int, len(mat[0]))
	for _, x := range a {
		pos := g[x]
		mx := make([]int, len(pos))
		for i, p := range pos {
			mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先把最大值算出来，再更新 rowMax 和 colMax
			ans = max(ans, mx[i])
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return
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
// [[3,1],[3,4]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,1,6],[-9,5,7]]\n
// @lcpr case=end

*/
