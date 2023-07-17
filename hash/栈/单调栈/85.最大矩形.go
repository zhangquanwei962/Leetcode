/*
 * @lc app=leetcode.cn id=85 lang=golang
 * @lcpr version=21909
 *
 * [85] 最大矩形
 */

// @lc code=start
// 该点为右下角
// O(mn) O(mn)
package main

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
		up := make([]int, m)
		down := make([]int, m)
		for i := range down {
			down[i] = m
		}
		stk := []int{-1} // 占位，避免判断,简化逻辑

		for i, l := range left {
			for len(stk) > 1 && left[stk[len(stk)-1]][j] >= l[j] {
				down[stk[len(stk)-1]] = i
				stk = stk[:len(stk)-1]
			}
			up[i] = stk[len(stk)-1] // 左边小于，右边小于等于
			stk = append(stk, i)
		}

		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			ans = max(ans, area)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0","0"]]\n
// @lcpr case=end

*/
