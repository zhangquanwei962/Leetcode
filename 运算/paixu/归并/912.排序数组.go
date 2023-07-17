/*
 * @lc app=leetcode.cn id=912 lang=golang
 * @lcpr version=21909
 *
 * [912] 排序数组
 */

// @lc code=start
package main

func sortArray(nums []int) []int {
	// rand.Seed(time.Now().UnixNano())
	return mergeSort(nums)
}
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums // 如果只有一个元素，则不需要排序，直接返回
	}
	// 分而治之：将数组分成两部分，递归调用归并排序
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// 合并有序子序列
	return merge(left, right)
}

// 合并两个有序数组
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// func quickSort(a []int, l, r int) {
// 	if l >= r {
// 		return
// 	}
// 	q := randomPartition(a, l, r)
// 	quickSort(a, l, q-1)
// 	quickSort(a, q+1, r)
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
// [5,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,1,2,0,0]\n
// @lcpr case=end

*/
