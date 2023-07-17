/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=21907
 *
 * [300] 最长递增子序列
 */

// @lc code=start
// f[i]表示以nums[i]结尾的最大长度
package main

import "sort"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func lengthOfLIS(nums []int) (res int) {
// 	f := make([]int, len(nums))
// 	for i, x := range nums {
// 		for j, y := range nums[:i] {
// 			if y < x {
// 				f[i] = max(f[i], f[j])
// 			}
// 		}
// 		f[i]++
// 	}
// 	for _, x := range f {
// 		res = max(res, x)
// 	}
// 	return
// }
func lengthOfLIS(nums []int) int {
	g := []int{}

	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) {
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)

}

// @lc code=end

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/
