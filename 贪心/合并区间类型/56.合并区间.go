/*
 * @lc app=leetcode.cn id=56 lang=golang
 * @lcpr version=21909
 *
 * [56] 合并区间
 */

// @lc code=start
package main

import (
	"sort"
)

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	l, maxR := intervals[0][0], intervals[0][1]

	for _, p := range intervals[1:] {
		if p[0] > maxR {
			ans = append(ans, []int{l, maxR})
			l = p[0]
		}
		maxR = max(maxR, p[1])
	}
	ans = append(ans, []int{l, maxR}) // 加入最后一个区间
	return
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
// [[1,3],[2,6],[8,10],[15,18]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[4,5]]\n
// @lcpr case=end

*/
