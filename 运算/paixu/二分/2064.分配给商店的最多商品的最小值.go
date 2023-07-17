/*
 * @lc app=leetcode.cn id=2064 lang=golang
 * @lcpr version=21908
 *
 * [2064] 分配给商店的最多商品的最小值
 */

// @lc code=start
// 最小化最大值 <=
package main

// func minimizedMaximum(n int, quantities []int) int {
// 	return sort.Search(1e5, func(max int) bool {
// 		cnt := 0
// 		for _, q := range quantities {
// 			cnt += (q + max) / (max + 1)
// 		}
// 		return cnt <= n
// 	}) + 1
// }

func minimizedMaximum(n int, quantities []int) int {
	max := 0
	for _, pile := range quantities {
		if pile > max {
			max = pile
		}
	}
	check := func(mid int) bool {
		time := 0
		for _, pile := range quantities {
			time += (pile + mid - 1) / mid
			// time += int(math.Ceil(float64(pile) / float64(mid)))
		}
		return time <= n
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
}

// @lc code=end

/*
// @lcpr case=start
// 6\n[11,6]\n
// @lcpr case=end

// @lcpr case=start
// 7\n[15,10,10]\n
// @lcpr case=end

// @lcpr case=start
// 1\n[100000]\n
// @lcpr case=end

*/
