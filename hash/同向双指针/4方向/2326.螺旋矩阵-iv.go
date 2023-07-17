/*
 * @lc app=leetcode.cn id=2326 lang=golang
 * @lcpr version=21909
 *
 * [2326] 螺旋矩阵 IV
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

var dirs = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func spiralMatrix(n int, m int, head *ListNode) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	for x, y, di := 0, 0, 0; head != nil; head = head.Next {
		ans[x][y] = head.Val
		d := dirs[di&3]
		if xx, yy := x+d.x, y+d.y; xx < 0 || xx >= n || yy < 0 || yy >= m || ans[xx][yy] != -1 {
			di++
			d = dirs[di&3]
		}
		x += d.x
		y += d.y
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 3\n5\n[3,0,2,6,8,1,7,9,4,2,5,5,0]\n
// @lcpr case=end

// @lcpr case=start
// 1\n4\n[0,1,2]\n
// @lcpr case=end

*/
