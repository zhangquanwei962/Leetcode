/*
 * @lc app=leetcode.cn id=930 lang=golang
 * @lcpr version=21908
 *
 * [930] 和相同的二元子数组
 */

// @lc code=start
// 先检查是否存在目标前缀和，再更新哈希表中的值
package main

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	cnt := map[int]int{0: 1}
	sum := 0
	for _, num := range nums {
		// cnt[sum]++
		sum += num
		cnt[sum]++
		ans += cnt[sum-goal]

	}
	// fmt.Println(cnt)
	return
}

// 标准做法
// func numSubarraysWithSum(nums []int, goal int) int {
// 	ans, cur := 0, 0
// 	first := make(map[int]int)
// 	first[0] = 1

// 	for _, v := range nums {
// 		cur += v
// 		if cnt, ok := first[cur-goal]; ok {
// 			ans += cnt
// 		}
// 		first[cur]++
// 	}

// 	return ans
// }

// func numSubarraysWithSum(nums []int, goal int) (ans int) {
// 	first := make(map[int]int)
// 	s := make([]int, len(nums)+1)
// 	for i := 0; i < len(nums); i++ {
// 		s[i+1] = s[i] + nums[i]
// 	}

// 	for i := 0; i < len(s); i++ {
// 		j, _ := first[s[i]-goal]
// 		ans += j
// 		first[s[i]]++
// 	}

// 	return
// }

// func numSubarraysWithSum(nums []int, goal int) (ans int) {
// 	preSum, s := make([]int, len(nums)+1), 0
// 	preSum[0] = 1

// 	for _, x := range nums {
// 		s += x
// 		if s >= goal {
// 			ans += preSum[s-goal]
// 		}
// 		preSum[s]++
// 	}
// 	return
// }

// @lc code=end

/*
// @lcpr case=start
// [1,0,1,0,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,0,0]\n0\n
// @lcpr case=end

*/
