/*
 * @lc app=leetcode.cn id=475 lang=golang
 * @lcpr version=21909
 *
 * [475] 供暖器
 */

// @lc code=start
// 最小加热半径的最大值，有单调性可以二分
// 每个房子都得找离他最近得heaters，然后就为这个房子的最短半径，然后统计所有房子的最大值就能得到最后的最小加热半径
package main

import (
	"math"
	"sort"
)

func findRadius(houses, heaters []int) (ans int) {
	sort.Ints(heaters)
	for _, house := range houses {
		j := sort.SearchInts(heaters, house+1)
		minDis := math.MaxInt32
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n[1,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,5]\n[2]\n
// @lcpr case=end

*/
