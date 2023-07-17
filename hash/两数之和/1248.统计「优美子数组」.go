/*
 * @lc app=leetcode.cn id=1248 lang=golang
 * @lcpr version=21908
 *
 * [1248] 统计「优美子数组」
 */

// @lc code=start
// 类似于两数之和

// map的话就是先记录再加
// 实时更新哈希表，以免出现i>=j的情况
// 其实就是先查询，在更新
package main

func numberOfSubarrays(nums []int, k int) (ans int) {
	cnt := map[int]int{0: 1}
	s := 0
	for _, x := range nums {
		s += x & 1
		cnt[s]++
		// if s >= k {
		ans += cnt[s-k]
		// }

	}
	return
}

// func numberOfSubarrays(nums []int, k int) (ans int) {
// 	preSum, s := make([]int, len(nums)+1), 0
// 	preSum[0] = 1

// 	for _, x := range nums {
// 		s += x & 1

// 		if s >= k {
// 			ans += preSum[s-k]
// 		}
// 		preSum[s]++
// 	}
// 	return
// }

// func numberOfSubarrays(nums []int, k int) (ans int) {
// 	n, n1 := 0, 0
// 	left, left1 := 0, 0

// 	for right, x := range nums {
// 		n += x & 1
// 		for n > k && left <= right { // 至多 k 个奇数
// 			n -= nums[left] & 1
// 			left++
// 		}

// 		n1 += x & 1
// 		for n1 > (k-1) && left1 <= right { // 至多 k-1 个奇数
// 			n1 -= nums[left1] & 1
// 			left1++
// 		}

// 		if n == k {
// 			ans += left1 - left
// 		}
// 	}
// 	return ans
// }

// @lc code=end

/*
// @lcpr case=start
// [1,1,2,1,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,4,6]\n1\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,1,2,2,1,2,2,2]\n2\n
// @lcpr case=end

*/
