/*
 * @lc app=leetcode.cn id=2513 lang=golang
 * @lcpr version=21909
 *
 * [2513] 最小化两个数组中的最大值
 */

// @lc code=start
// 最小化最大值
// 答案越大，越能组成
package main

func minimizeSet(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	lcm := d1 / gcd(d1, d2) * d2
	// a := []int{1, 2, 2, 4, 4, 4, 5, 6, 7, 7, 8}
	// i := sort.Search(9, func(k int) bool {
	// 	extra := 0
	// 	for i := len(a) - 1; i > 0; i-- {
	// 		extra = max(a[i]+extra-k, 0)
	// 	}
	// 	fmt.Println(a[0]+extra, k, a[0]+extra <= k)
	// 	return a[0]+extra <= k
	// })
	// fmt.Println(i)
	// return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
	// 	left1 := max(uniqueCnt1-x/d2+x/lcm, 0)
	// 	left2 := max(uniqueCnt2-x/d1+x/lcm, 0)
	// 	common := x - x/d1 - x/d2 + x/lcm
	// 	return common >= left1+left2
	// })
	check := func(mid int) bool {
		left1 := max(uniqueCnt1-mid/d2+mid/lcm, 0)
		left2 := max(uniqueCnt2-mid/d1+mid/lcm, 0)
		common := mid - mid/d1 - mid/d2 + mid/lcm
		return common >= left1+left2
	}
	left, right := 1, (uniqueCnt1+uniqueCnt2)*2
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

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// 2\n7\n1\n3\n
// @lcpr case=end

// @lcpr case=start
// 3\n5\n2\n1\n
// @lcpr case=end

// @lcpr case=start
// 2\n4\n8\n2\n
// @lcpr case=end

*/
