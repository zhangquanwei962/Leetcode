/*
 * @lc app=leetcode.cn id=1857 lang=golang
 * @lcpr version=21909
 *
 * [1857] 有向图中最大颜色值
 */

// @lc code=start
package main

func largestPathValue(colors string, edges [][]int) (ans int) {
	n := len(colors)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		if v == w {
			return -1
		}
		g[v] = append(g[v], w)
		deg[w]++
	}
	cnt := 0
	dp := make([][26]int, n)
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		cnt++
		dp[v][colors[v]-'a']++
		ans = max(ans, dp[v][colors[v]-'a'])
		for _, w := range g[v] {
			for i, c := range dp[v] {
				dp[w][i] = max(dp[w][i], c)
			}
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	if cnt < n {
		return -1
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
// "abaca"\n[[0,1],[0,2],[2,3],[3,4]]\n
// @lcpr case=end

// @lcpr case=start
// "a"\n[[0,0]]\n
// @lcpr case=end

*/
