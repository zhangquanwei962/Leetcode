/*
 * @lc app=leetcode.cn id=2709 lang=golang
 * @lcpr version=21909
 *
 * [2709] 最大公约数遍历
 */

// @lc code=start
// gcd实际是翻滚求余，而得到数一定是某个大家都有质因数的倍数
// 质数或者质因数
package main

import "fmt"

const mxa = 100000 + 1

var p = [mxa][]int{}

func init() {
	for i := 2; i < mxa; i++ {
		if len(p[i]) == 0 {
			for j := i; j < mxa; j += i {
				p[j] = append(p[j], i)
			}
		}
	}
	fmt.Println(p)
}
func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	mx := 0
	for _, x := range nums {
		if mx < x {
			mx = x
		}
	}
	fa := make([]int, mx+n+1)
	for i := range fa {
		fa[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx != fy {
			fa[fy] = fx
		}
	}

	for i, x := range nums {
		for _, j := range p[x] {
			union(i, n+j) // 把质因数映射到n+1~n+max上了，n+p就是质数p对应的点
		}

	}
	// 类似于set
	has := make(map[int]struct{})
	for i := 0; i < n; i++ {
		x := find(i)
		has[x] = struct{}{}
	}

	if len(has) == 1 {
		return true
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6]\n
// @lcpr case=end

// @lcpr case=start
// [3,9,5]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,12,8]\n
// @lcpr case=end

*/
