/*
 * @lc app=leetcode.cn id=1691 lang=golang
 * @lcpr version=21909
 *
 * [1691] 堆叠长方体的最大高度
 */

// @lc code=start
// O(n^2) O(n)
package main

import "sort"

func maxHeight(cuboids [][]int) (ans int) {
	for _, c := range cuboids {
		sort.Ints(c)
	}

	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]
		return a[0] < b[0] || a[0] == b[0] && (a[1] < b[1] || a[1] == b[1] && a[2] < b[2])
	})

	f := make([]int, len(cuboids))
	for i, c2 := range cuboids {
		for j, c1 := range cuboids[:i] {
			if c1[1] <= c2[1] && c1[2] <= c2[2] { // 排序后，c1[0] <= c2[0] 恒成立
				f[i] = max(f[i], f[j]) // c1 可以堆在 c2 上
			}
		}
		f[i] += c2[2]
		ans = max(ans, f[i])
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
// [[50,45,20],[95,37,53],[45,23,12]]\n
// @lcpr case=end

// @lcpr case=start
// [[38,25,45],[76,35,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]]\n
// @lcpr case=end

*/
