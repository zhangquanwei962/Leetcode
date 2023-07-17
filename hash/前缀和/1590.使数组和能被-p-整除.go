/*
 * @lc app=leetcode.cn id=1590 lang=golang
 * @lcpr version=21909
 *
 * [1590] 使数组和能被 P 整除
 */

// @lc code=start
package main

func minSubarray(nums []int, p int) int {
	x := 0
	for _, v := range nums {
		x += v
	}
	x %= p
	if x == 0 {
		return 0 // 移除空子数组（这个 if 可以不要）
	}

	n := len(nums)
	ans, s := n, 0
	// 由于下面 i 是从 0 开始的，前缀和下标就要从 -1 开始了
	last := map[int]int{s: -1}
	for i, v := range nums {
		s += v
		last[s%p] = i
		if j, ok := last[(s-x+p)%p]; ok {
			ans = min(ans, i-j)
		}

	}
	if ans < n {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,4,2]\n6\n
// @lcpr case=end

// @lcpr case=start
// [6,3,5,2]\n9\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n7\n
// @lcpr case=end

// @lcpr case=start
// [1000000000,1000000000,1000000000]\n3\n
// @lcpr case=end

*/
