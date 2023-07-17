/*
 * @lc app=leetcode.cn id=剑指 Offer 51 lang=golang
 * @lcpr version=21909
 *
 * [剑指 Offer 51] 数组中的逆序对
 */

// @lc code=start
// O(nlogn) O(n)
package main

import (
	"fmt"
	"sort"
)

func reversePairs(nums []int) int {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", nums[i])
	}

	bit := BIT{
		n:    n,
		tree: make([]int, n+1),
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += bit.query(nums[i] - 1)
		bit.update(nums[i])
	}
	return ans
}

type BIT struct {
	n    int
	tree []int
}

func (b BIT) lowbit(x int) int { return x & (-x) }

func (b BIT) query(x int) int {
	ret := 0
	for x > 0 {
		ret += b.tree[x]
		x -= b.lowbit(x)
	}
	return ret
}

func (b BIT) update(x int) {
	for x <= b.n {
		b.tree[x]++
		x += b.lowbit(x)
	}
}

// var a []int

// func reversePairs(nums []int) (ans int) {
// 	discretization(nums)
// 	n := len(nums)
// 	bit := BIT{
// 		n:    n,
// 		tree: make([]int, n+1),
// 	}

// 	for i := n - 1; i >= 0; i-- {
// 		id := getId(nums[i])
// 		ans += bit.query(id - 1)
// 		bit.update(id)
// 	}
// 	return
// }

// func discretization(nums []int) {
// 	set := map[int]struct{}{}
// 	for _, num := range nums {
// 		set[num] = struct{}{}
// 	}
// 	a = make([]int, 0, len(nums))
// 	for num := range set {
// 		a = append(a, num)
// 	}
// 	sort.Ints(a)
// }
// func getId(x int) int {
// 	return sort.SearchInts(a, x) + 1
// }

// type BIT struct {
// 	n    int   // length
// 	tree []int // array
// }

// func (b BIT) lowbit(x int) int { return x & (-x) }

// func (b BIT) query(x int) int {
// 	ret := 0
// 	for x > 0 {
// 		ret += b.tree[x]
// 		x -= b.lowbit(x)
// 	}
// 	return ret
// }

// func (b BIT) update(x int) {
// 	for x <= b.n {
// 		b.tree[x]++
// 		x += b.lowbit(x)
// 	}
// }

// @lc code=end

/*
// @lcpr case=start
// [7,5,6,4]\n
// @lcpr case=end

*/
