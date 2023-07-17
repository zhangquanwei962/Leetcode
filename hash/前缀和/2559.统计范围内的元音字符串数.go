/*
 * @lc app=leetcode.cn id=2559 lang=golang
 * @lcpr version=21908
 *
 * [2559] 统计范围内的元音字符串数
 */

// @lc code=start
package main

import "strings"

func vowelStrings(words []string, queries [][]int) []int {
	sum := make([]int, len(words)+1)
	for i, w := range words {
		sum[i+1] = sum[i]
		if strings.Contains("aeiou", w[:1]) && strings.Contains("aeiou", w[len(w)-1:]) {
			sum[i+1]++
		}

	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sum[q[1]+1] - sum[q[0]]
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// ["aba","bcb","ece","aa","e"]\n[[0,2],[1,4],[1,1]]\n
// @lcpr case=end

// @lcpr case=start
// ["a","e","i"]\n[[0,2],[0,1],[2,2]]\n
// @lcpr case=end

*/
