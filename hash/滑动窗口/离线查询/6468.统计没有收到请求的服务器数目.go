/*
 * @lc app=leetcode.cn id=6468 lang=golang
 * @lcpr version=21909
 *
 * [6468] 统计没有收到请求的服务器数目
 */

// @lc code=start
// 离线查询+滑动窗口
// 对于询问，为了不打乱顺序，可以创建一个下标数组对其排序。
// O(n+mlogm+qlogq) O(n+q)
package main

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	type pair struct{ q, i int }
	qs := make([]pair, len(queries))
	for i, q := range queries {
		qs[i] = pair{q, i}
	}

	sort.Slice(qs, func(i, j int) bool { return qs[i].q < qs[j].q })
	sort.Slice(logs, func(i, j int) bool { return logs[i][1] < logs[j][1] })

	ans := make([]int, len(queries))
	cnt := make([]int, n+1)

	outOfRange, left, right := n, 0, 0
	for _, p := range qs {
		for ; right < len(logs) && logs[right][1] <= p.q; right++ { // 进入窗口
			i := logs[right][0]
			if cnt[i] == 0 {
				outOfRange--
			}
			cnt[i]++
		}
		for ; left < len(logs) && logs[left][1] < p.q-x; left++ { // 离开窗口
			i := logs[left][0]
			cnt[i]--
			if cnt[i] == 0 {
				outOfRange++
			}
		}
		ans[p.i] = outOfRange
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 3\n[[1,3],[2,6],[1,5]]\n5\n[10,11]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[[2,4],[2,1],[1,2],[3,1]]\n2\n[3,4]\n
// @lcpr case=end

*/
