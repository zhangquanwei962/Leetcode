/*
 * @lc app=leetcode.cn id=327 lang=golang
 * @lcpr version=21909
 *
 * [327] 区间和的个数
 */

// @lc code=start
package main

import "sort"

type fenwick struct {
	tree []int
}

func (f fenwick) inc(i int) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i]++
	}
}

func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

func (f fenwick) query(l, r int) (res int) {
	return f.sum(r) - f.sum(l-1)
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
	n := len(nums)

	// 计算前缀和 preSum，以及后面统计时会用到的所有数字 allNums
	allNums := make([]int, 1, 3*n+1)
	preSum := make([]int, n+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
		allNums = append(allNums, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}

	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i <= 3*n; i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}

	// 遍历 preSum，利用树状数组计算每个前缀和对应的合法区间数
	t := fenwick{make([]int, k+1)}
	t.inc(kth[0])
	for _, sum := range preSum[1:] {
		left, right := kth[sum-upper], kth[sum-lower]
		cnt += t.query(left, right)
		t.inc(kth[sum])
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [-2,5,-1]\n-2\n2\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n0\n
// @lcpr case=end

*/
