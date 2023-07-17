/*
 * @lc app=leetcode.cn id=1024 lang=golang
 * @lcpr version=21909
 *
 * [1024] 视频拼接
 */

// @lc code=start
// O(n) O(n)
package main

func videoStitching(clips [][]int, time int) int {
	maxn := make([]int, time+1)
	for _, p := range clips {
		start, end := p[0], p[1]
		if start < time {
			maxn[start] = max(maxn[start], end)
		}
	}
	ans, cur_right, next_right := 0, 0, 0
	for i, t := range maxn[:time] {
		next_right = max(next_right, t)
		if i == cur_right {
			if i == next_right {
				return -1
			}
			cur_right = next_right
			ans++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]\n10\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]]\n9\n
// @lcpr case=end

*/
