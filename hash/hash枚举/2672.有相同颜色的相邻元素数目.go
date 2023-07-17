/*
 * @lc app=leetcode.cn id=2672 lang=golang
 * @lcpr version=21909
 *
 * [2672] 有相同颜色的相邻元素数目
 */

// @lc code=start
package main

func colorTheArray(n int, queries [][]int) []int {
	ans := make([]int, len(queries))
	a := make([]int, n+2) // 避免越界讨论
	cnt := 0

	for qi, q := range queries {
		i, c := q[0]+1, q[1]
		if a[i] > 0 {
			if a[i] == a[i-1] {
				cnt--
			}
			if a[i] == a[i+1] {
				cnt--
			}
		}

		a[i] = c
		if a[i] == a[i-1] {
			cnt++
		}
		if a[i] == a[i+1] {
			cnt++
		}
		ans[qi] = cnt
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 4\n[[0,2],[1,2],[3,1],[1,1],[2,1]]\n
// @lcpr case=end

// @lcpr case=start
// 1\n[[0,100000]]\n
// @lcpr case=end

*/
