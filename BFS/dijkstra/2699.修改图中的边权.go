/*
 * @lc app=leetcode.cn id=2699 lang=golang
 * @lcpr version=21908
 *
 * [2699] 修改图中的边权
 */

// @lc code=start
package main

import "math"

func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
	type edge struct{ to, eid int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, i})
		g[y] = append(g[y], edge{x, i}) // 建图，额外记录边的编号
	}

	var delta int
	dis := make([][2]int, n)
	for i := range dis {
		dis[i] = [2]int{math.MaxInt, math.MaxInt}
	}
	dis[source] = [2]int{}
	dijkstra := func(k int) { // 这里 k 表示第一次/第二次
		vis := make([]bool, n)
		for {
			// 找到当前最短路，去更新它的邻居的最短路
			// 根据数学归纳法，dis[x][k] 一定是最短路长度
			x := -1
			for y, b := range vis {
				if !b && (x < 0 || dis[y][k] < dis[x][k]) {
					x = y
				}
			}
			if x == destination { // 起点 source 到终点 destination 的最短路已确定
				return
			}
			vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
			for _, e := range g[x] {
				y, wt := e.to, edges[e.eid][2]
				if wt == -1 {
					wt = 1 // -1 改成 1
				}
				if k == 1 && edges[e.eid][2] == -1 {
					// 第二次 Dijkstra，改成 w
					w := delta + dis[y][0] - dis[x][1]
					if w > wt {
						wt = w
						edges[e.eid][2] = w // 直接在 edges 上修改
					}
				}
				// 	for (int j = 1; j <= n; j ++ )//用t去更新其他边
				// dist[j] = min(dist[j], dist[t] + g[t][j]);
				// 更新最短路
				dis[y][k] = min(dis[y][k], dis[x][k]+wt)
			}
		}
	}

	dijkstra(0)
	delta = target - dis[destination][0]
	if delta < 0 { // -1 全改为 1 时，最短路比 target 还大
		return nil
	}

	dijkstra(1)
	if dis[destination][1] < target { // 最短路无法再变大，无法达到 target
		return nil
	}

	for _, e := range edges {
		if e[2] == -1 { // 剩余没修改的边全部改成 1
			e[2] = 1
		}
	}
	return edges
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// 5\n[[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]]\n0\n1\n5\n
// @lcpr case=end

// @lcpr case=start
// 3\n[[0,1,-1],[0,2,5]]\n0\n2\n6\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[1,0,4],[1,2,3],[2,3,5],[0,3,-1]]\n0\n2\n6\n
// @lcpr case=end

*/
