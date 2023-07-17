/*
 * @lc app=leetcode.cn id=2517 lang=golang
 * @lcpr version=21908
 *
 * [2517] 礼盒的最大甜蜜度
 */

// @lc code=start
// 最大化最小值 >=
// 当return left的时候，>=和<可以互换
package main

import (
	"sort"
)

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)

	var check func(int) bool
	check = func(d int) bool {
		cnt, pre := 1, price[0]
		for _, p := range price {
			if p >= pre+d {
				cnt++
				pre = p
			}

		}
		return cnt >= k

	}
	var binary func(int, int) int
	binary = func(left, right int) int {
		for left+1 < right {
			mid := left + (right-left)/2
			if check(mid) {
				left = mid
			} else {
				right = mid
			}
		}
		return left
	}

	return binary(0, (price[len(price)-1]-price[0])/(k-1)+1)

}

// @lc code=end

/*
// @lcpr case=start
// [13,5,1,8,21,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,3,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7]\n2\n
// @lcpr case=end

*/
