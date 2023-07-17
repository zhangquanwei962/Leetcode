/*
 * @lc app=leetcode.cn id=2258 lang=golang
 * @lcpr version=21909
 *
 * [2258] 逃离火灾
 */

// @lc code=start
// sort.Search()是第一个满足f(i)==true
// 那么f(i)-1就是最后一个满足f(i)==false
package main

import "sort"

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := sort.Search(m*n+1, func(t int) bool {
		fire := make([][]bool, m)
		for i := range fire {
			fire[i] = make([]bool, n)
		}
		f := []pair{}
		for i, row := range grid {
			for j, v := range row {
				if v == 1 {
					fire[i][j] = true
					f = append(f, pair{i, j})
				}
			}
		}
		spreadFire := func() {
			tmp := f
			f = nil
			for _, p := range tmp {
				for _, d := range dirs {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] == 0 {
						fire[x][y] = true
						f = append(f, pair{x, y})
					}
				}
			}
		}
		for ; t > 0 && len(f) > 0; t-- {
			spreadFire() // 蔓延至多 t 分钟的火势
		}
		if fire[0][0] { // 起点着火，寄
			return true
		}

		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		q := []pair{{}}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, p := range tmp {
				if !fire[p.x][p.y] {
					for _, d := range dirs {
						if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && !fire[x][y] && grid[x][y] == 0 {
							if x == m-1 && y == n-1 { // 我们安全了…暂时。
								return false
							}
							vis[x][y] = true
							q = append(q, pair{x, y})
						}
					}
				}
			}
			spreadFire() // 蔓延 1 分钟的火势
		}
		return true // 寄
	}) - 1
	if ans < m*n {
		return ans
	}
	return 1e9
}

// @lc code=end

/*
// @lcpr case=start
// [[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0,0],[0,1,2,0],[0,2,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0],[2,2,0],[1,2,0]]\n
// @lcpr case=end

*/
