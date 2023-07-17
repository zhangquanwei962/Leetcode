/*
 * @lc app=leetcode.cn id=1971 lang=golang
 * @lcpr version=21909
 *
 * [1971] 寻找图中是否存在路径
 */

// @lc code=start
package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}

	for _, e := range edges {
		p[find(e[0])] = p[find(e[1])]
	}
	return find(source) == find(destination)
}

// @lc code=end

/*
// @lcpr case=start
// 3\n[[0,1],[1,2],[2,0]]\n0\n2\n
// @lcpr case=end

// @lcpr case=start
// 6\n[[0,1],[0,2],[3,5],[5,4],[4,3]]\n0\n5\n
// @lcpr case=end

*/
