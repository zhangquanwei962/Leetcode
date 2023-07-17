/*
 * @lc app=leetcode.cn id=2368 lang=golang
 * @lcpr version=21907
 *
 * [2368] 受限条件下可到达节点的数目
 */

// @lc code=start
package main

func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
	r := make(map[int]bool, len(restricted))
	for _, x := range restricted {
		r[x] = true
	}

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		if !r[x] && !r[y] {
			g[x] = append(g[x], y)
			g[y] = append(g[y], x)
		}
	}

	var dfs func(int, int)
	dfs = func(x, fa int) {
		ans++
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	return
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[[0,1],[1,2],[3,1],[4,0],[0,5],[5,6]]\n[4,5]\n
// @lcpr case=end

// @lcpr case=start
// 7\n[[0,1],[0,2],[0,5],[0,4],[3,2],[6,5]]\n[4,2,1]\n
// @lcpr case=end

*/
