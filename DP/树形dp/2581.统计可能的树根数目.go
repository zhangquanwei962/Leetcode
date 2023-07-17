/*
 * @lc app=leetcode.cn id=2581 lang=golang
 * @lcpr version=21909
 *
 * [2581] 统计可能的树根数目
 */

// @lc code=start
// 只有[x,y] [y,x]这俩的正确性被改变了
package main

func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // 建图
	}

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根时，猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[1,2],[1,3],[4,2]]\n[[1,3],[0,1],[1,0],[2,4]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2],[2,3],[3,4]]\n[[1,0],[3,4],[2,1],[3,2]]\n1\n
// @lcpr case=end

*/
