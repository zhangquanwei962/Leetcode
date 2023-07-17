/*
 * @lc app=leetcode.cn id=2050 lang=golang
 * @lcpr version=21909
 *
 * [2050] 并行课程 III
 */

// @lc code=start
package main

func minimumTime(n int, relations [][]int, time []int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range relations {
		v, w := e[0]-1, e[1]-1
		g[v] = append(g[v], w)
		deg[w]++
	}

	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}

	f := make([]int, n)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		f[v] += time[v]
		ans = max(ans, f[v])
		for _, w := range g[v] {
			f[w] = max(f[v], f[w]) //更新f[w]
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
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
// 3\n[[1,3],[2,3]]\n[3,2,5]\n
// @lcpr case=end

// @lcpr case=start
// 5\n[[1,5],[2,5],[3,5],[3,4],[4,5]]\n[1,2,3,4,5]\n
// @lcpr case=end

*/
