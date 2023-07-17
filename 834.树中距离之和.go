/*
 * @lc app=leetcode.cn id=834 lang=golang
 * @lcpr version=21909
 *
 * [834] 树中距离之和
 */

// @lc code=start
// 换根dp O(n) O(n)
package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := make([]int, n)
	size := make([]int, n)

	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		ans[0] += depth
		size[x] = 1
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, depth+1)
				size[x] += size[y]
			}
		}
	}

	dfs(0, -1, 0)

	var reRoot func(int, int)
	reRoot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				ans[y] = ans[x] + n - 2*size[y]
				reRoot(y, x)
			}
		}
	}
	reRoot(0, -1)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 6\n[[0,1],[0,2],[2,3],[2,4],[2,5]]\n
// @lcpr case=end

// @lcpr case=start
// 1\n[]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

*/
