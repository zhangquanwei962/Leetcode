/*
 * @lc app=leetcode.cn id=873 lang=golang
 * @lcpr version=21909
 *
 * [873] 最长的斐波那契子序列的长度
 */

// @lc code=start
// 当值域都是[0,1,...,n-1]可以置换
// 当整个值域都是单调递增时候可以置换，（离散化） 10^9
package main

func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	p := make(map[int]int, n)

	for i, x := range arr {
		p[x] = i
	}
	// f[i][j]是以X_{i}\X_{j}结尾的长度
	// 类似于等差数列
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	for i, x := range arr {
		for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
			if k, ok := p[x-arr[j]]; ok {
				f[j][i] = max(f[k][j]+1, 3)
				ans = max(ans, f[j][i])
			}
		}
	}
	return
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
// [1,2,3,4,5,6,7,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,7,11,12,14,18]\n
// @lcpr case=end

*/
