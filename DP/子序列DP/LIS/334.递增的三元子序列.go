/*
 * @lc app=leetcode.cn id=334 lang=golang
 * @lcpr version=21909
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
package main

import "sort"

func increasingTriplet(nums []int) bool {
	g := []int{}

	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) {
			g = append(g, x)
			if len(g) == 3 {
				return true
			}
		} else {
			g[j] = x
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,5,0,4,6]\n
// @lcpr case=end

*/
