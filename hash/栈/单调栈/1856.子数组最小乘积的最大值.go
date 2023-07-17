/*
 * @lc app=leetcode.cn id=1856 lang=golang
 * @lcpr version=21909
 *
 * [1856] 子数组最小乘积的最大值
 */

// @lc code=start
// 首先声明这个区间都是(1,4)==[2,3]
// 前缀和是s[right]-s[left]=[left,right)

// left是小于 right是大于等于
package main

func maxSumMinProduct(nums []int) int {
	const mod int = 1e9 + 7
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	pre := make([]int, n+1)

	for i, x := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] >= x {
			right[st[len(st)-1]] = i // i 恰好是栈顶的右边界
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1] + 1
		st = append(st, i)
	}

	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}

	ans := 0
	for i, x := range nums {
		ans = max(ans, (pre[right[i]]-pre[left[i]])*x)
	}
	return ans % mod

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
// [1,2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,3,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,1,5,6,4,2]\n
// @lcpr case=end

*/
