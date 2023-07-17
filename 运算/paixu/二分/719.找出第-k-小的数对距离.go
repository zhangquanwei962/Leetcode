/*
 * @lc app=leetcode.cn id=719 lang=golang
 * @lcpr version=21908
 *
 * [719] 找出第 K 小的数对距离
 */

// @lc code=start
package main

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	var check func(int) bool
	check = func(mid int) bool {
		cnt, i := 0, 0
		for j, num := range nums {
			for i < j && num-nums[i] > mid {
				i++
			}
			cnt += j - i
		}
		return cnt >= k
	}

	var bina func(int, int) int
	bina = func(left, right int) int {
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

	return bina(-1, nums[len(nums)-1]-nums[0]+1)
}

// return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
// 	cnt, i := 0, 0
// 	for j, num := range nums {
// 		for num-nums[i] > mid {
// 			i++
// 		}
// 		cnt += j - i
// 	}
// 	return cnt >= k
// })

// @lc code=end

/*
// @lcpr case=start
// [1,3,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,6,1]\n3\n
// @lcpr case=end

*/
