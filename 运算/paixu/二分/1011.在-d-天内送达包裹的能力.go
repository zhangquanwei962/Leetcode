/*
 * @lc app=leetcode.cn id=1011 lang=golang
 * @lcpr version=21909
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
// 希望查询满足条件的最小值！
// >=就是在有等于基础上的最大值（最大化最小值 left）
// <=就是在有等于基础上的最小值 （最小化最大值）
// 裸二分不是，罗二分>=就是第一个等于的
package main

func shipWithinDays(weights []int, days int) int {
	// 确定二分查找左右边界
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	check := func(mid int) bool {
		day, sum := 1, 0
		for _, w := range weights {
			if sum+w > mid {
				day++
				sum = 0
			}
			sum += w
		}
		return day > days
	}

	binary := func(left, right int) int {
		for left+1 < right {
			mid := left + (right-left)>>1
			// fmt.Println(left, right, mid)
			if check(mid) {
				left = mid
			} else {
				right = mid
			}
		}
		return right
	}
	x := binary(left-1, right+1)

	return x

	// return left + sort.Search(right-left, func(x int) bool {
	// 	x += left
	// 	day := 1 // 需要运送的天数
	// 	sum := 0 // 当前这一天已经运送的包裹重量之和
	// 	for _, w := range weights {
	// 		if sum+w > x {
	// 			day++
	// 			sum = 0
	// 		}
	// 		sum += w
	// 	}
	// 	return day <= days
	// })
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10]\n5\n
// @lcpr case=end

// @lcpr case=start
// [3,2,2,4,1,4]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,1]\n4\n
// @lcpr case=end

*/
