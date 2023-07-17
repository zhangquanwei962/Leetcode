/*
 * @lc app=leetcode.cn id=841 lang=golang
 * @lcpr version=21909
 *
 * [841] 钥匙和房间
 */

// @lc code=start
package main

var (
	num int
	vis []bool
)

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	vis = make([]bool, n)
	dfs(rooms, 0)
	return num == n
}

func dfs(rooms [][]int, x int) {
	vis[x] = true
	num++
	for _, it := range rooms[x] {
		if !vis[it] {
			dfs(rooms, it)
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[1],[2],[3],[]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[3,0,1],[2],[0]]\n
// @lcpr case=end

*/
