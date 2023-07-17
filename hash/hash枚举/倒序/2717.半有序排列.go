/*
 * @lc app=leetcode.cn id=2717 lang=golang
 * @lcpr version=21909
 *
 * [2717] 半有序排列
 */

// @lc code=start
package main

func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	var p, q int
	for i, x := range nums {
		if x == 1 {
			p = i
		} else if x == n {
			q = i
		}
	}
	if p < q {
		return p + n - 1 - q
	}
	return p + n - 2 - q // 1 向左移动的时候和 n 交换了一次
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,4,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,4,2,5]\n
// @lcpr case=end

*/
