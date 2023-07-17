/*
 * @lc app=leetcode.cn id=2555 lang=golang
 * @lcpr version=21908
 *
 * [2555] 两个线段获得的最多奖品
 */

// @lc code=start
package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// pre[right+1]代表线段右端点不超过pos[right]能覆盖多少奖品
func maximizeWin(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	pre := make([]int, n+1)
	left := 0

	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1)
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2,2,3,3,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n0\n
// @lcpr case=end

*/
