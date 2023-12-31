/*
 * @lc app=leetcode.cn id=1579 lang=golang
 * @lcpr version=21909
 *
 * [1579] 保证图可完全遍历
 */

// @lc code=start
// 可以删除的最大，等价于需要最少去联通（保留最少的边）
// 抽象为两个并查集在保持图连通，
// 即各自的连通分量均为一的前提下，可以删去的最大边数量
package main

type unionFind struct {
	parent, size []int
	setCount     int // 当前连通分量数目
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	uf.setCount--
	return true
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	ans := len(edges)
	alice, bob := newUnionFind(n), newUnionFind(n)
	for _, e := range edges {
		x, y := e[1]-1, e[2]-1
		if e[0] == 3 && (!alice.inSameSet(x, y) || !bob.inSameSet(x, y)) {
			// 保留这条公共边
			alice.union(x, y)
			bob.union(x, y)
			ans--
		}
	}
	uf := [2]*unionFind{alice, bob}
	for _, e := range edges {
		if tp := e[0]; tp < 3 && uf[tp-1].union(e[1]-1, e[2]-1) {
			// 保留这条独占边
			ans--
		}
	}
	if alice.setCount > 1 || bob.setCount > 1 {
		return -1
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 4\n[[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[3,1,2],[3,2,3],[1,1,4],[2,1,4]]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[3,2,3],[1,1,2],[2,3,4]]\n
// @lcpr case=end

*/
