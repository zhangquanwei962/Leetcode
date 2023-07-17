/*
 * @lc app=leetcode.cn id=1494 lang=golang
 * @lcpr version=21909
 *
 * [1494] 并行课程 II
 */

// @lc code=start
package main

import (
	"math"
	"math/bits"
)

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	pre1 := make([]int, n)
	for _, r := range relations {
		pre1[r[1]-1] |= 1 << (r[0] - 1) // r[1] 的先修课程集合，下标改从 0 开始
	}

	u := 1<<n - 1 // 全集
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i == 0 { // 空集
			return
		}
		p := &memo[i]
		if *p != -1 { // 之前算过了
			return *p
		}
		defer func() { *p = res }() // 记忆化
		ci := u ^ i                 // i 的补集
		i1 := 0
		for j, p := range pre1 {
			if i>>j&1 == 1 && p|ci == ci { // p 在 i 的补集中，可以学（否则这学期一定不能学）j属于i
				i1 |= 1 << j
			}
		}
		if bits.OnesCount(uint(i1)) <= k { // 如果个数小于 k，则可以全部学习，不再枚举子集
			return dfs(i^i1) + 1
		}
		res = math.MaxInt
		for j := i1; j > 0; j = (j - 1) & i1 { // 枚举 i1 的子集 j
			if bits.OnesCount(uint(j)) == k {
				res = min(res, dfs(i^j)+1)
			}
		}
		return res
	}
	return dfs(u)
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
// 4\n[[2,1],[3,1],[1,4]]\n2\n
// @lcpr case=end

// @lcpr case=start
// 5\n[[2,1],[3,1],[4,1],[1,5]]\n2\n
// @lcpr case=end

// @lcpr case=start
// 11\n[]\n2\n
// @lcpr case=end

*/
