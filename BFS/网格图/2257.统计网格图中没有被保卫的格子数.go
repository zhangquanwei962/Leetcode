/*
 * @lc app=leetcode.cn id=2257 lang=golang
 * @lcpr version=21909
 *
 * [2257] 统计网格图中没有被保卫的格子数
 */

// @lc code=start
// 负数的与一直是-1
package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) (ans int) {
	type pair struct{ x, y, k int }
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	q := []pair{}

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	// x := (-2 >> 0) & 1
	// fmt.Println(strconv.FormatInt(int64(x), 2))
	for _, guard := range guards {
		x, y := guard[0], guard[1]
		grid[x][y] = -1 // 记录是警卫
		for k := 0; k < 4; k++ {
			q = append(q, pair{x, y, k})
		}
	}

	for _, wall := range walls {
		x, y := wall[0], wall[1]
		grid[x][y] = -2 // 记录的是墙
	}

	for len(q) > 0 {
		x, y, k := q[0].x, q[0].y, q[0].k
		q = q[1:]
		nx, ny := x+dx[k], y+dy[k]
		if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] >= 0 {
			// 沿着视线方向的下一个坐标合法，且不为警卫或墙
			if (grid[nx][ny]>>k)&1 == 0 { // 不在集合
				grid[nx][ny] |= 1 << k // 加入
				q = append(q, pair{nx, ny, k})
			}
		}
	}

	for _, row := range grid {
		for _, x := range row {
			if x == 0 {
				ans++
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 4\n6\n[[0,0],[1,1],[2,3]]\n[[0,1],[2,2],[1,4]]\n
// @lcpr case=end

// @lcpr case=start
// 3\n3\n[[1,1]]\n[[0,1],[1,0],[2,1],[1,2]]\n
// @lcpr case=end

*/
