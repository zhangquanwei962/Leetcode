/*
 * @lc app=leetcode.cn id=1713 lang=golang
 * @lcpr version=21909
 *
 * [1713] 得到子序列的最少操作次数
 */

// @lc code=start
// O(mlogm+n) O(m+n)
// 置换LIS O(mn)会超时，不会影响答案个数和相互顺序
// 提示：要是target变成[0,1...n-1]就会简单不少， 枚举arr
// 注意到target单调递增，arr也要单调递增，变为求arr的单调递增序列
package main

import "sort"

func minOperations(target []int, arr []int) int {
	n := len(target)
	p := make(map[int]int, n)
	for i, v := range target {
		p[v] = i
	}

	d := []int{}
	for _, x := range arr {
		if idx, has := p[x]; has {
			if j := sort.SearchInts(d, idx); j < len(d) {
				d[j] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}

// @lc code=end

/*
// @lcpr case=start
// [5,1,3]\n[9,4,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [6,4,8,1,3,2]\n[4,7,6,2,3,8,6,1]\n
// @lcpr case=end

*/
