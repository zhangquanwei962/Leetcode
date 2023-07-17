/*
 * @lc app=leetcode.cn id=2104 lang=golang
 * @lcpr version=21909
 *
 * [2104] 子数组范围和
 */

// @lc code=start
package main

func solve(nums []int) (ans int64) {
	n := len(nums)
	left := make([]int, n)  // left[i] 为左侧严格大于 num[i] 的最近元素位置（不存在时为 -1）
	right := make([]int, n) // right[i] 为右侧大于等于 num[i] 的最近元素位置（不存在时为 n）
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	for i, v := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] <= v {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	for i, v := range nums {
		ans += int64(i-left[i]) * int64(right[i]-i) * int64(v)
	}
	return
}

func subArrayRanges(nums []int) int64 {
	ans := solve(nums)
	for i, v := range nums { // 小技巧：所有元素取反后算的就是最小值的贡献
		nums[i] = -v
	}
	return ans + solve(nums)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,-2,-3,4,1]\n
// @lcpr case=end

*/
