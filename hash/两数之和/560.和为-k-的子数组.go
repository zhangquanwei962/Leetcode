/*
 * @lc app=leetcode.cn id=560 lang=golang
 * @lcpr version=21908
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
package main

func subarraySum(nums []int, k int) (ans int) {
	first := make(map[int]int)
	s := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		s[i+1] = s[i] + nums[i]
	}

	for i := 0; i < len(s); i++ {
		j, _ := first[s[i]-k]
		ans += j
		first[s[i]]++
	}
	return
}

// 先更新在查询和先查询在更新都一样对于这个
// 	cnt := map[int]int{}
// 	sum := 0
// 	for _, x := range nums {
// 		cnt[sum]++
// 		sum += x
// 		ans += cnt[sum-k]
// 	}
// 	return
// }

// @lc code=end

/*
// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

*/
