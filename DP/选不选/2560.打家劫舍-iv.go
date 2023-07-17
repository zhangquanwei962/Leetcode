/*
 * @lc app=leetcode.cn id=2560 lang=golang
 * @lcpr version=21909
 *
 * [2560] 打家劫舍 IV
 */

// @lc code=start
// 最小化最大值，返回的一定是right
package main

func minCapability(nums []int, k int) int {
	check := func(mid int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x <= mid {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	}
	binary := func(left, right int) int {
		for left+1 < right {
			mid := left + (right-left)>>1
			if check(mid) {
				right = mid
			} else {
				left = mid
			}
		}
		return right
	}
	left, right := -1, 1000000000+1
	return binary(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,5,9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n2\n
// @lcpr case=end

*/
