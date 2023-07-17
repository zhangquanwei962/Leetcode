/*
 * @lc app=leetcode.cn id=215 lang=golang
 * @lcpr version=21909
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
package main

import (
	"math/rand"
	"time"
)

// func findKthLargest(nums []int, k int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
// }

// func quickSelect(a []int, l, r, index int) int {
// 	q := randomPartition(a, l, r)
// 	if q == index {
// 		return a[index]
// 	} else if q < index {
// 		return quickSelect(a, q+1, r, index)
// 	}
// 	return quickSelect(a, l, q-1, index)
// }

// func randomPartition(a []int, l, r int) int {
// 	i := rand.Int()%(r-l+1) + l
// 	a[i], a[r] = a[r], a[i]
// 	return partition(a, l, r)
// }

// func partition(a []int, l, r int) int {
// 	x := a[r]
// 	i := l - 1
// 	for j := i; j < r; j++ {
// 		if a[j] <= x {
// 			i++
// 			a[i], a[j] = a[j], a[i]
// 		}
// 	}
// 	a[i+1], a[r] = a[r], a[i+1]
// 	return i + 1
// }

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, len(nums)-k)
}
func quickSelect(a []int, k int) int {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > l && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
}

// func quickSelect(a []int, l, r, index int) int {
// 	q := randomPartition(a, l, r)
// 	if q == index {
// 		return a[q]
// 	} else if q < index {
// 		return quickSelect(a, q+1, r, index)
// 	}
// 	return quickSelect(a, l, q-1, index)
// }

// func randomPartition(a []int, l, r int) int {
// 	i := rand.Int()%(r-l+1) + l
// 	a[i], a[r] = a[r], a[i]
// 	return partition(a, l, r)
// }

// func partition(a []int, l, r int) int {
// 	x := a[r]
// 	i := l - 1
// 	for j := l; j < r; j++ {
// 		if a[j] <= x {
// 			i++
// 			a[i], a[j] = a[j], a[i]
// 		}
// 	}
// 	a[i+1], a[r] = a[r], a[i+1]
// 	return i + 1
// }

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/
