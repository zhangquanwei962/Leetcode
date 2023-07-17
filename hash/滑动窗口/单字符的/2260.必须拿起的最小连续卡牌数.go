/*
 * @lc app=leetcode.cn id=2260 lang=golang
 * @lcpr version=21909
 *
 * [2260] 必须拿起的最小连续卡牌数
 */

// @lc code=start
package main

import (
	"math"
)

func minimumCardPickup(cards []int) int {
	left, ans := 0, math.MaxInt
	cnt := make(map[int]int)
	for right, x := range cards {
		cnt[x]++
		for cnt[x] > 1 {
			ans = min(ans, right-left+1)
			cnt[cards[left]]--
			if cnt[cards[left]] == 0 {
				delete(cnt, cards[left])
			}
			left++
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
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
// [3,4,2,3,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,5,3]\n
// @lcpr case=end

*/
