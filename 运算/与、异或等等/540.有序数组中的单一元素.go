/*
 * @lc app=leetcode.cn id=540 lang=golang
 * @lcpr version=21909
 *
 * [540] 有序数组中的单一元素
 */

// @lc code=start
// i^1可以是i+1 可以是i-1
package main

func singleNonDuplicate(a []int) int {
	// i := sort.Search(len(a)-1, func(i int) bool {
	// 	return a[i] != a[i^1]
	// })
	left, right := -1, len(a)
	for left+1 < right {
		mid := left + (right-left)/2
		if a[mid] != a[mid^1] {
			right = mid
		} else {
			left = mid
		}
	}
	return a[right]
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2,3,3,4,4,8,8]\n
// @lcpr case=end

// @lcpr case=start
// [3,3,7,7,10,11,11]\n
// @lcpr case=end

*/
