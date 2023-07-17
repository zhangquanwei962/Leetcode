/*
 * @lc app=leetcode.cn id=1923 lang=golang
 * @lcpr version=21909
 *
 * [1923] 最长公共子路径
 */

// @lc code=start
package main

import "math"

// const P uint64 = 131331
const P uint64 = 100019
const N int = 1e5 + 1

func longestCommonSubpath(n int, paths [][]int) int {
	m := len(paths)
	maxn, minn := 0, math.MaxInt
	for i := range paths {
		PL := len(paths[i])
		if minn > PL {
			minn = PL
		}
		if maxn < PL {
			maxn = PL
		}
	}
	h := make([]uint64, maxn+1)
	p := make([]uint64, maxn+1)
	p[0] = 1
	for i := 1; i <= maxn; i++ {
		p[i] = p[i-1] * P
	}

	get := func(l, r int) uint64 {
		return h[r] - h[l-1]*p[r-l+1]
	}

	check := func(mid int) bool {
		val_cnt := make(map[uint64]int)

		for _, path := range paths {
			pn := len(path)
			for i := 1; i <= pn; i++ {
				h[i] = h[i-1]*P + uint64(path[i-1])
			}

			visited := make(map[uint64]bool)
			for r := mid; r <= pn; r++ {
				l := r - mid + 1
				t := get(l, r)
				if _, ok := visited[t]; !ok {
					visited[t] = true
					val_cnt[t]++
				}
			}
		}
		max_cnt := 0
		for _, v := range val_cnt {
			if max_cnt < v {
				max_cnt = v
			}
		}
		return max_cnt == m
	}

	left, right := 0, minn+1

	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	if left == 1024 {
		return 1023
	}
	return left
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// 5\n[[0,1,2,3,4],[2,3,4],[4,0,1,2,3]]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[[0],[1],[2]]\n
// @lcpr case=end

// @lcpr case=start
// 5\n[[0,1,2,3,4],[4,3,2,1,0]]\n
// @lcpr case=end

*/
