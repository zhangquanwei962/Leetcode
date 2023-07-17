/*
 * @lc app=leetcode.cn id=1686 lang=golang
 * @lcpr version=21909
 *
 * [1686] 石子游戏 VI
 */

// @lc code=start
// 归纳法证明
// a1 a2 b1 b2 c1=a1-b2 c2=a2-b1 c=a1+b1-(a2+b2) ? 0
package main

import (
	"sort"
)

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	arr := make([][2]int, n)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		arr[i] = [2]int{aliceValues[i] + bobValues[i], i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			a += aliceValues[arr[i][1]]
		} else {
			b += bobValues[arr[i][1]]
		}
	}
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n[2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,3]\n[1,6,7]\n
// @lcpr case=end

*/
