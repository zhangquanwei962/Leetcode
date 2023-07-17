/*
 * @lc app=leetcode.cn id=189 lang=golang
 * @lcpr version=21908
 *
 * [189] 轮转数组
 */

// @lc code=start
// 双指针 交换，i+n-i-1=n-1

// 分组交换
package main

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// func reverse(a []int) {
// 	for i, n := 0, len(a); i < n/2; i++ {
// 		a[i], a[n-i-1] = a[n-i-1], a[i]
// 	}
// }
// func rotate(nums []int, k int) {
// 	k %= len(nums)
// 	reverse(nums)
// 	reverse(nums[:k])
// 	reverse(nums[k:])
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [-1,-100,3,99]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end
*/
