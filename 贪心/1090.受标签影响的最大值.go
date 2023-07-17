/*
 * @lc app=leetcode.cn id=1090 lang=golang
 * @lcpr version=21908
 *
 * [1090] 受标签影响的最大值
 */

// @lc code=start
package main

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}

	sort.Slice(idx, func(i, j int) bool {
		return values[idx[i]] > values[idx[j]]
	})

	ans, choose := 0, 0
	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		label := labels[idx[i]]

		if cnt[label] == useLimit {
			continue
		}

		choose++
		ans += values[idx[i]]
		cnt[label]++
		if choose == numWanted {
			break
		}
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [5,4,3,2,1]\n[1,1,2,2,3]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,4,3,2,1]\n[1,3,3,3,2]\n3\n2\n
// @lcpr case=end

// @lcpr case=start
// [9,8,8,7,6]\n[0,0,0,1,1]\n3\n1\n
// @lcpr case=end

*/
