/*
 * @lc app=leetcode.cn id=1889 lang=golang
 * @lcpr version=21909
 *
 * [1889] 装包裹的最小浪费空间
 */

// @lc code=start
// 由于包裹的尺寸之和是固定的，因此目标转换为最小化箱子的尺寸之和。
// 将包裹按尺寸排序后，按照箱子尺寸划分包裹，使得每个包裹都装入能装该包裹的最小的箱子。这种贪心策略能使箱子的尺寸总和尽可能地小。
package main

import (
	"math"
	"sort"
)

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	ans := math.MaxInt64

	for _, box := range boxes {
		sort.Ints(box)
		if box[len(box)-1] < packages[len(packages)-1] { // 最大的箱子不够装最大的包裹
			continue
		}

		res, l := 0, 0
		for _, v := range box {
			// 划分包裹：当前箱子 v 可以装入下标在 [l, r) 区间内的包裹
			r := sort.SearchInts(packages, v+1)
			res += (r - l) * v // 统计箱子尺寸之和
			l = r
		}
		ans = min(ans, res)
	}
	if ans < math.MaxInt64 {
		for _, v := range packages {
			ans -= v // 减去每个包裹尺寸
		}
		return ans % (1e9 + 7)
	}
	return -1
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
// [2,3,5]\n[[4,8],[2,8]]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n[[1,4],[2,3],[3,4]]\n
// @lcpr case=end

// @lcpr case=start
// [3,5,8,10,11,12]\n[[12],[11,9],[10,5,14]]\n
// @lcpr case=end

*/
