/*
 * @lc app=leetcode.cn id=1986 lang=golang
 * @lcpr version=21909
 *
 * [1986] 完成任务的最少工作时间段
 */

// @lc code=start
package main

func minSessions(tasks []int, sessionTime int) (ans int) {
	n := len(tasks)
	m := 1 << n
	// 预处理所有子集的子集和，复杂度 O(1+2+4+...+2^(n-1)) = O(2^n)
	sum := make([]int, m)
	for i, t := range tasks {
		for j, k := 0, 1<<i; j < k; j++ {
			sum[j|k] = sum[j] + t
		}
	}
	f := make([]int, m)
	for i := range f {
		f[i] = n
	}
	f[0] = 0
	for s := range f {
		// 枚举 s 的所有子集 sub，若 sub 耗时不超过 sessionTime，则将 f[s^sub]+1 转移到 f[s] 上
		for sub := s; sub > 0; sub = (sub - 1) & s {
			if sum[sub] <= sessionTime && f[s^sub]+1 < f[s] {
				f[s] = f[s^sub] + 1
			}
		}
	}
	return f[m-1]
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [3,1,3,1,1]\n8\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n15\n
// @lcpr case=end

*/
