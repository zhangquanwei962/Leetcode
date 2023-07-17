/*
 * @lc app=leetcode.cn id=2432 lang=golang
 * @lcpr version=21907
 *
 * [2432] 处理用时最长的那个任务的员工
 */

// @lc code=start
package main

func hardestWorker(n int, logs [][]int) (ans int) {
	var mx, last int

	for _, log := range logs {
		uid, t := log[0], log[1]
		t -= last
		if mx < t || (mx == t && uid < ans) {
			mx = t
			ans = uid
		}
		last += t
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 10\n[[0,3],[2,5],[0,9],[1,15]]\n
// @lcpr case=end

// @lcpr case=start
// 26\n[[1,1],[3,7],[2,12],[7,17]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[0,10],[1,20]]\n
// @lcpr case=end

*/
