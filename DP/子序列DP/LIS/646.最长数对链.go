/*
 * @lc app=leetcode.cn id=646 lang=golang
 * @lcpr version=21909
 *
 * [646] 最长数对链
 */

// @lc code=start
// 无重复区间个数
package main

import "sort"

func findLongestChain(points [][]int) int {
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

// func findLongestChain(intervals [][]int) int {
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
// 		// 找的是最小值
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
// [[1,2], [2,3], [3,4]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[7,8],[4,5]]\n
// @lcpr case=end

*/
