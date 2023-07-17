/*
 * @lc app=leetcode.cn id=1377 lang=golang
 * @lcpr version=21908
 *
 * [1377] T 秒后青蛙的位置
 */

// @lc code=start
package main

func frogPosition(n int, edges [][]int, t int, target int) (ans float64) {

	g := make([][]int, n+1)

	g[1] = []int{0}

	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int, int, int) bool

	dfs = func(x, fa, leftT, prod int) bool {
		// t 秒后必须在 target（恰好到达，或者 target 是叶子停在原地）
		if x == target && (leftT == 0 || len(g[x]) == 1) {
			ans = 1 / float64(prod)
			return true
		}

		if x == target || leftT == 0 {
			return false
		}

		for _, y := range g[x] {
			if y != fa && dfs(y, x, leftT-1, prod*(len(g[x])-1)) {
				return true
			}
		}

		return false

	}

	dfs(1, 0, t, 1)
	return

	// g := make([][]int, n+1)

	// for _, e := range edges {
	// 	x, y := e[0], e[1]
	// 	g[x] = append(g[x], y)
	// 	g[y] = append(g[y], x)
	// }

	// vis := make([]bool, n+1)

	// var dfs func(int, int) float64

	// dfs = func(i int, t int) float64 {
	// 	nxt := len(g[i])

	// 	if i > 1 {
	// 		nxt -= 1
	// 	}

	// 	if t == 0 || nxt == 0 {
	// 		if i == target {
	// 			return 1.0
	// 		} else {
	// 			return 0.0
	// 		}
	// 	}

	// 	vis[i] = true
	// 	ans := 0.0

	// 	for _, j := range g[i] {
	// 		if !vis[j] {
	// 			ans += dfs(j, t-1)
	// 		}
	// 	}

	// 	return ans / float64(nxt)

	// }

	// return dfs(1, t)
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// 7\n[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]\n1\n7\n
// @lcpr case=end

*/
