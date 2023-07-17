/*
 * @lc app=leetcode.cn id=410 lang=golang
 * @lcpr version=21909
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
// 最小化最大值,是<=
// 最大化最小值，是>=
package main

func splitArray(nums []int, m int) int {
	left, right := 0, 0
	for _, num := range nums {
		right += num // 右边界，所有元素的和
		if left < num {
			left = num // 左边界，数组中最大的元素
		}
	}
	var check func(int) bool
	check = func(mid int) bool {
		sum, count := 0, 1
		for _, x := range nums {
			if sum+x > mid {
				count++
				sum = x
			} else {
				sum += x
			}
		}
		return count > m
	}
	left, right = left-1, right+1

	for left+1 < right {
		mid := left + (right-left)/2 // 二分左边界

		if check(mid) { // 左边界太大，需要把左边界调大，扩大子数组的平均值
			left = mid
		} else { // 左边界太小，需要把左边界调小，缩小子数组的平均值
			right = mid
		}
	}

	return right
}

// package main

// func splitArray(nums []int, m int) int {
// 	left, right := 0, 0
// 	for _, num := range nums {
// 		right += num // 右边界，所有元素的和
// 		if left < num {
// 			left = num // 左边界，数组中最大的元素
// 		}
// 	}
// 	var check func(int) bool
// 	check = func(mid int) bool {
// 		sum, count := 0, 1
// 		for _, x := range nums {
// 			if sum+x > mid {
// 				count++
// 				sum = x
// 			} else {
// 				sum += x
// 			}
// 		}
// 		return count <= m
// 	}
// 	left, right = left-1, right+1

// 	for left+1 < right {
// 		mid := left + (right-left)/2 // 二分左边界

// 		if check(mid) {
// 			right = mid
// 		} else {
// 			left = mid
// 		}
// 	}

// 	return right
// }
// @lc code=end

/*
// @lcpr case=start
// [7,2,5,10,8]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,4]\n3\n
// @lcpr case=end

*/
