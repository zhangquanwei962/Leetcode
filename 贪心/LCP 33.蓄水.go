/*
 * @lc app=leetcode.cn id=LCP 33 lang=golang
 * @lcpr version=21908
 *
 * [LCP 33] 蓄水
 */

// @lc code=start
package main

import "math"

func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	maxk := 0
	for _, v := range vat {
		maxk = max(maxk, v)
	}
	if maxk == 0 {
		return 0
	}
	ans := math.MaxInt32
	for k := 1; k <= maxk && k < ans; k++ {
		cur := k
		for i := 0; i < n; i++ {
			least := vat[i]/k + btoi(vat[i]%k != 0)
			cur += max(least-bucket[i], 0)
		}
		ans = min(ans, cur)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n[6,8]\n
// @lcpr case=end

// @lcpr case=start
// [9,0,1]\n[0,2,2]\n
// @lcpr case=end

*/
