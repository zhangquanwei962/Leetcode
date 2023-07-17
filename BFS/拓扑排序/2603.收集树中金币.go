/*
 * @lc app=leetcode.cn id=2603 lang=golang
 * @lcpr version=21909
 *
 * [2603] 收集树中金币
 */

// @lc code=start
package main

func collectTheCoins(coins []int, edges [][]int) (ans int) {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
		deg[x]++
		deg[y]++
	}

	// 用拓扑排序「剪枝」：去掉没有金币的子树
	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 1 && coins[i] == 0 { // 无金币叶子
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 && coins[y] == 0 {
				q = append(q, y)
			}
		}
	}

	// 再次拓扑排序
	for i, d := range deg {
		if d == 1 && coins[i] == 1 { // 有金币叶子
			q = append(q, i)
		}
	}
	if len(q) <= 1 { // 至多一个有金币的叶子，直接收集
		return
	}
	time := make([]int, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 {
				time[y] = time[x] + 1 // 记录入队时间
				q = append(q, y)
			}
		}
	}

	// 统计答案
	for _, e := range edges {
		if time[e[0]] >= 2 && time[e[1]] >= 2 {
			ans += 2
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,0,0,0,1]\n[[0,1],[1,2],[2,3],[3,4],[4,5]]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,1,1,0,0,1]\n[[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5,7]]\n
// @lcpr case=end

*/
