/*
 * @lc app=leetcode.cn id=1042 lang=golang
 * @lcpr version=21908
 *
 * [1042] 不邻接植花
 */

// @lc code=start
package main

func gardenNoAdj(n int, paths [][]int) []int {
	g := make([][]int, n)
	for _, e := range paths {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	color := make([]int, n)

	for i, nodes := range g {
		used := [5]bool{}
		for _, j := range nodes {
			used[color[j]] = true
		}

		for color[i]++; used[color[i]]; color[i]++ {

		}
	}
	return color
}

// @lc code=end

/*
// @lcpr case=start
// 3\n[[1,2],[2,3],[3,1]]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[1,2],[3,4]]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]\n
// @lcpr case=end

*/
