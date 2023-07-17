/*
 * @lc app=leetcode.cn id=1039 lang=golang
 * @lcpr version=21909
 *
 * [1039] 多边形三角剖分的最低得分
 */

// @lc code=start
package main

import "math"

func minScoreTriangulation(v []int) int {
	n := len(v)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示还没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i+1 == j { // 只有两个点，无法组成三角形
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res := math.MaxInt
		for k := i + 1; k < j; k++ { // 枚举顶点 k
			res = min(res, dfs(i, k)+dfs(k, j)+v[i]*v[j]*v[k])
		}
		*p = res
		return res
	}
	return dfs(0, n-1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// func minScoreTriangulation(v []int) int {
//     n := len(v)
//     f := make([][]int, n)
//     for i := range f {
//         f[i] = make([]int, n)
//     }
//     for i := n - 3; i >= 0; i-- {
//         for j := i + 2; j < n; j++ {
//             f[i][j] = math.MaxInt
//             for k := i + 1; k < j; k++ {
//                 f[i][j] = min(f[i][j], f[i][k]+f[k][j]+v[i]*v[j]*v[k])
//             }
//         }
//     }
//     return f[0][n-1]
// }

// func min(a, b int) int { if b < a { return b }; return a }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [3,7,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,1,4,1,5]\n
// @lcpr case=end

*/
