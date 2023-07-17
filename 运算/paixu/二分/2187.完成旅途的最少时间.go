/*
 * @lc app=leetcode.cn id=2187 lang=golang
 * @lcpr version=21908
 *
 * [2187] 完成旅途的最少时间
 */

// @lc code=start
package main

func minimumTime(time []int, totalTrips int) int64 {
	max := 0
	for _, x := range time {
		if max < x {
			max = x
		}
	}

	var check func(int) bool
	check = func(mid int) bool {
		total := 0
		for _, t := range time {
			total += mid / t
		}
		return total >= totalTrips
	}

	binary := func(left, right int) int {
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
	return int64(binary(0, totalTrips*max+1))
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/
