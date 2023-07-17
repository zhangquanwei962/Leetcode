/*
 * @lc app=leetcode.cn id=1743 lang=golang
 * @lcpr version=21909
 *
 * [1743] 从相邻元素对还原数组
 */

// @lc code=start
// 中间数字有俩相邻，开头结尾就一个，答案呼之欲出
// O(n) O(n)
package main

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	g := make(map[int][]int, n) // 哈希表记录
	for _, p := range adjacentPairs {
		v, w := p[0], p[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n)
	for e, adj := range g { // 确定开头
		if len(adj) == 1 {
			ans[0] = e
			break
		}
	}

	ans[1] = g[ans[0]][0]
	for i := 2; i < n; i++ {
		adj := g[ans[i-1]] // 顺序不定
		if ans[i-2] == adj[0] {
			ans[i] = adj[1]
		} else {
			ans[i] = adj[0]
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1],[3,4],[3,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[4,-2],[1,4],[-3,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[100000,-100000]]\n
// @lcpr case=end

*/
