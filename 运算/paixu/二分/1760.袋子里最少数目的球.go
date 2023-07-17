/*
 * @lc app=leetcode.cn id=1760 lang=golang
 * @lcpr version=21908
 *
 * [1760] 袋子里最少数目的球
 */

// @lc code=start
package main

func minimumSize(nums []int, maxOperations int) int {
	var check func(int) bool
	check = func(mid int) bool {
		s := 0
		for _, v := range nums {
			s += (v - 1) / mid
		}
		return s <= maxOperations
	}
	var binary func(int, int) int
	binary = func(left, right int) int {
		for left+1 < right {
			mid := left + (right-left)/2
			if check(mid) {
				right = mid
			} else {
				left = mid
			}
		}
		return right
	}
	return binary(0, 1e9+1)
}

// @lc code=end

/*
// @lcpr case=start
// [9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [2,4,8,2]\n4\n
// @lcpr case=end

// @lcpr case=start
// [7,17]\n2\n
// @lcpr case=end

*/
