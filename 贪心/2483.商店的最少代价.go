/*
 * @lc app=leetcode.cn id=2483 lang=golang
 * @lcpr version=21909
 *
 * [2483] 商店的最少代价
 */

// @lc code=start
// 贪心
package main

import "strings"

func bestClosingTime(customers string) (ans int) {
	cost := strings.Count(customers, "Y")
	mincost := cost
	for i, c := range customers {
		if c == 'N' {
			cost++
		} else {
			cost--
			if mincost > cost {
				mincost = cost
				ans = i + 1
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "YYNY"\n
// @lcpr case=end

// @lcpr case=start
// "NNNNN"\n
// @lcpr case=end

// @lcpr case=start
// "YYYY"\n
// @lcpr case=end

*/
