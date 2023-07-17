/*
 * @lc app=leetcode.cn id=452 lang=golang
 * @lcpr version=21909
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
// 不重叠区间个数 -> 覆盖区间
package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}

// func findMinArrowShots(intervals [][]int) int {
// 	n := len(intervals)
// 	if n == 0 {
// 		return 0
// 	}
// 	sort.Slice(intervals, func(i, j int) bool {
// 		a, b := intervals[i], intervals[j]
// 		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
// 	})

// 	f := []int{}
// 	for _, e := range intervals {
// 		j := sort.Search(len(f), func(k int) bool {
// 			return f[k] >= e[0]
// 		})
// 		if j == len(f) {
// 			f = append(f, e[1])
// 		} else if f[j] > e[1] {
// 			f[j] = e[1]
// 		}

// 	}
// 	return len(f)
// }

// @lc code=end

/*
// @lcpr case=start
// [[10,16],[2,8],[1,6],[7,12]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,4],[5,6],[7,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[2,3],[3,4],[4,5]]\n
// @lcpr case=end

*/
