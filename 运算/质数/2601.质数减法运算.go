/*
 * @lc app=leetcode.cn id=2601 lang=golang
 * @lcpr version=21909
 *
 * [2601] 质数减法运算
 */

// @lc code=start
// 质数筛
package main

import (
	"fmt"
	"sort"
)

var p = []int{0} // 哨兵，避免二分越界

func init() {
	const mx = 1000
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			p = append(p, i) // 预处理质数列表
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
	fmt.Println(p)
}

func primeSubOperation(nums []int) bool {
	pre := 0
	for _, x := range nums {
		if x <= pre {
			return false
		}
		pre = x - p[sort.SearchInts(p, x-pre)-1] // 减去 < x-pre 的最大质数
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [4,9,6,10]\n
// @lcpr case=end

// @lcpr case=start
// [6,8,11,12]\n
// @lcpr case=end

// @lcpr case=start
// [5,8,3]\n
// @lcpr case=end

*/
