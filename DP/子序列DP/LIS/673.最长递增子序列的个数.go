/*
 * @lc app=leetcode.cn id=673 lang=golang
 * @lcpr version=21909
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
package main

import "sort"

func findNumberOfLIS(nums []int) int {
	d := [][]int{}
	cnt := [][]int{}
	for _, v := range nums {
		i := sort.Search(len(d), func(i int) bool { return d[i][len(d[i])-1] >= v })
		c := 1
		if i > 0 {
			k := sort.Search(len(d[i-1]), func(k int) bool { return d[i-1][k] < v })
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}

		if i == len(d) {
			d = append(d, []int{v})
			cnt = append(cnt, []int{0, c})
		} else {
			d[i] = append(d[i], v)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1]+c)
		}
	}
	c := cnt[len(cnt)-1]
	return c[len(c)-1]
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n
// @lcpr case=end

*/
