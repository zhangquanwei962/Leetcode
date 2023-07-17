/*
 * @lc app=leetcode.cn id=1000 lang=golang
 * @lcpr version=21909
 *
 * [1000] 合并石头的最低成本
 */

// @lc code=start
// 保证区间非空，不是非法
package main

import "math"

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

	s := make([]int, n+1)
	for i, x := range stones {
		s[i+1] = s[i] + x // 前缀和
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示还没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == j { // 只有一堆石头，无需合并
			return
		}
		ptr := &memo[i][j]
		if *ptr != -1 {
			return *ptr
		}
		defer func() { *ptr = res }()
		res = math.MaxInt
		for m := i; m < j; m += k - 1 {
			res = min(res, dfs(i, m)+dfs(m+1, j))
		}
		if (j-i)%(k-1) == 0 { // 可以合并成一堆
			res += s[j+1] - s[i]
		}
		return
	}
	return dfs(0, n-1)
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
// [3,2,4,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [3,5,1,2,6]\n3\n
// @lcpr case=end

*/
