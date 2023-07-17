/*
 * @lc app=leetcode.cn id=2455 lang=golang
 * @lcpr version=21908
 *
 * [2455] 可被三整除的偶数的平均值
 */

// @lc code=start
package main

func averageValue(nums []int) int {
	total := 0
	k := 0
	for _, a := range nums {
		if a%6 == 0 {
			total += a
			k++
		}
	}
	if k > 0 {
		return total / k
	} else {
		return 0
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,6,10,12,15]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,4,7,10]\n
// @lcpr case=end

*/
