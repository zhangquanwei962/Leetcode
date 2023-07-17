/*
 * @lc app=leetcode.cn id=2611 lang=golang
 * @lcpr version=21909
 *
 * [2611] 老鼠和奶酪
 */

// @lc code=start
package main

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x
		reward1[i] -= x
	}

	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, x := range reward1[:k] {
		ans += x
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,3,4]\n[4,4,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n[1,1]\n2\n
// @lcpr case=end

*/
