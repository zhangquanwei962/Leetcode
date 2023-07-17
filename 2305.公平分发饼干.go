/*
 * @lc app=leetcode.cn id=2305 lang=golang
 * @lcpr version=21909
 *
 * [2305] 公平分发饼干
 */

// @lc code=start
// 定义 f[i][j] 表示前i个孩子分配的饼干集合为 j 时，前
// i 个孩子的不公平程度的最小值。

// f[i][j]=min_{s$j}max(f[i-1][j\s],sums[s])
// 类比01背包消去第一维
// O(k*3^n) O(2^n)
package main

func distributeCookies(a []int, k int) int {
	n := 1 << len(a)
	sum := make([]int, n)
	for i, v := range a {
		for j, bit := 0, 1<<i; j < bit; j++ {
			sum[bit|j] = sum[j] + v
		}
	}

	f := append([]int{}, sum...)
	for i := 1; i < k; i++ {
		for j := n - 1; j > 0; j-- {
			for s := j; s > 0; s = (s - 1) & j { // 枚举子集
				f[j] = min(f[j], max(f[j^s], sum[s]))
			}
		}
	}
	return f[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [8,15,10,20,8]\n2\n
// @lcpr case=end

// @lcpr case=start
// [6,1,3,2,2,4,1,2]\n3\n
// @lcpr case=end

*/
