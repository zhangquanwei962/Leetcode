/*
 * @lc app=leetcode.cn id=875 lang=golang
 * @lcpr version=21908
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
package main

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	check := func(mid int) bool {
		time := 0
		for _, pile := range piles {
			time += (pile + mid - 1) / mid
			// time += int(math.Ceil(float64(pile) / float64(mid)))
		}
		return time <= h
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
	return binary(0, max)
	// return 1 + sort.Search(max-1, func(speed int) bool {
	// 	speed++
	// 	time := 0
	// 	for _, pile := range piles {
	// 		time += (pile + speed - 1) / speed
	// 	}
	// 	return time <= h
	// })
}

// @lc code=end

/*
// @lcpr case=start
// [3,6,7,11]\n8\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n5\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n6\n
// @lcpr case=end

*/
