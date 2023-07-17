/*
 * @lc app=leetcode.cn id=1655 lang=golang
 * @lcpr version=21909
 *
 * [1655] 分配重复整数
 */

// @lc code=start
package main

func canDistribute(nums []int, quantity []int) bool {
	// 预处理 quantity 每个子集的子集和
	m := 1 << len(quantity)
	sum := make([]int, m)
	for i, v := range quantity {
		bit := 1 << i
		for j := 0; j < bit; j++ {
			sum[bit|j] = sum[j] + v
		}
	}

	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}

	n := len(cnt)
	// cnt数组的前i个元素，能否满足顾客子集合j的订单数目
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m)
		f[i][0] = true
	}
	i := 0
	for _, c := range cnt {
		for j, ok := range f[i] {
			if ok {
				f[i+1][j] = true
				continue
			}
			for sub := j; sub > 0; sub = (sub - 1) & j { // 枚举 j 的子集 sub
				if sum[sub] <= c && f[i][j^sub] { // 判断这 c 个数能否全部分给 sub，并且除了 sub 以外的 j 中的顾客也满足
					f[i+1][j] = true
					break
				}
			}
		}
		i++
	}
	return f[n][m-1]

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,2]\n[2,2]\n
// @lcpr case=end

*/
