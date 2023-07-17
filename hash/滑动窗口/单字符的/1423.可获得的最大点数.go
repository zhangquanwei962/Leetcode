/*
 * @lc app=leetcode.cn id=1423 lang=golang
 * @lcpr version=21909
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
package main

import "math"

func maxScore(cardPoints []int, k int) int {
	sum := 0
	ans := math.MaxInt
	left, s, cnt := 0, 0, 0
	n := len(cardPoints)

	for _, x := range cardPoints {
		s += x
		sum += x
		cnt++
		if cnt >= n-k {
			ans = min(ans, s)
			s -= cardPoints[left]
			left++
			cnt--
		}
	}
	if n == k {
		return sum
	}
	return sum - ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2]\n2\n
// @lcpr case=end

// @lcpr case=start
// [9,7,7,9,7,7,9]\n7\n
// @lcpr case=end

// @lcpr case=start
// [1,1000,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,79,80,1,1,1,200,1]\n3\n
// @lcpr case=end

*/
