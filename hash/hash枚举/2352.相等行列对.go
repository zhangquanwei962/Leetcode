/*
 * @lc app=leetcode.cn id=2352 lang=golang
 * @lcpr version=21908
 *
 * [2352] 相等行列对
 */

// @lc code=start
package main

func equalPairs(grid [][]int) (ans int) {
	cnt := map[[200]int]int{}
	for _, row := range grid {
		a := [200]int{}
		for j, v := range row {
			a[j] = v
		}
		cnt[a]++
	}

	for j := range grid[0] {
		a := [200]int{}
		for i, row := range grid {
			a[i] = row[j]
		}
		ans += cnt[a]
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [[3,2,1],[1,7,6],[2,7,7]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]\n
// @lcpr case=end

*/
