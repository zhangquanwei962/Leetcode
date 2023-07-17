/*
 * @lc app=leetcode.cn id=2578 lang=golang
 * @lcpr version=21909
 *
 * [2578] 最小和分割
 */

// @lc code=start
// 贪心
package main

import (
	"sort"
	"strconv"
)

func splitNum(num int) int {
	s := []byte(strconv.Itoa(num))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	a := [2]int{}
	for i, c := range s {
		a[i%2] = a[i%2]*10 + int(c-'0') // 按照奇偶下标分组
	}
	return a[0] + a[1]
}

// @lc code=end

/*
// @lcpr case=start
// 4325\n
// @lcpr case=end

// @lcpr case=start
// 687\n
// @lcpr case=end

*/
