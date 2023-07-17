/*
 * @lc app=leetcode.cn id=LCP 41 lang=golang
 * @lcpr version=21909
 *
 * [LCP 41] 黑白翻转棋
 */

// @lc code=start
package main

var dir8 = []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func cntFlip(g [][]byte) (cnt int) {
	n, m := len(g), len(g[0])
	for { // 不断循环，直到无法翻转白棋
		flip := false
		for i, row := range g {
			for j, ch := range row {
				if ch != 'O' {
					continue
				}
				// 对于每个白棋，检查行、列或对角线中是否存在两端均有黑棋，且中间没有空白
			outer:
				for k, d := range dir8[:4] {
					// 检查一个方向
					for x, y := i, j; ; x, y = x+d.x, y+d.y {
						if x < 0 || x >= n || y < 0 || y >= m || g[x][y] == '.' {
							continue outer
						}
						if g[x][y] == 'X' {
							break
						}
					}
					// 检查相反的另一方向
					d = dir8[k+4]
					for x, y := i, j; ; x, y = x+d.x, y+d.y {
						if x < 0 || x >= n || y < 0 || y >= m || g[x][y] == '.' {
							continue outer
						}
						if g[x][y] == 'X' {
							break
						}
					}
					// 将其翻转为黑棋
					g[i][j] = 'X'
					flip = true
					cnt++
					break
				}
			}
		}
		if !flip {
			return
		}
	}
}

func flipChess(g []string) (ans int) {
	g2 := make([][]byte, len(g))
	for i, row := range g {
		for j, ch := range row {
			if ch == '.' {
				for k, r := range g {
					g2[k] = []byte(r)
				}
				g2[i][j] = 'X' // 放上黑棋
				if cnt := cntFlip(g2); cnt > ans {
					ans = cnt
				}
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// ["....X.","....X.","XOOO..","......","......"]`>\n
// @lcpr case=end

// @lcpr case=start
// [".X.",".O.","XO."]`>\n
// @lcpr case=end

// @lcpr case=start
// [".......",".......",".......","X......",".O.....","..O....","....OOX"]`>\n
// @lcpr case=end

*/
