/*
 * @lc app=leetcode.cn id=1283 lang=golang
 * @lcpr version=21908
 *
 * [1283] 使结果不超过阈值的最小除数
 */

// @lc code=start
// 和严格大于的，一定不是解
package main

func smallestDivisor(nums []int, threshold int) int {
	max := 0
	for _, pile := range nums {
		if pile > max {
			max = pile
		}
	}

	var check func(int) bool
	check = func(mid int) bool {
		total := 0
		for _, num := range nums {
			total += (num + mid - 1) / mid
		}
		return total <= threshold
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
	return binary(0, max)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,5,9]\n6\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5,7,11]\n11\n
// @lcpr case=end

// @lcpr case=start
// [19]\n5\n
// @lcpr case=end

*/
