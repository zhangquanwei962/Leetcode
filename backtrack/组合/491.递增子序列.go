/*
 * @lc app=leetcode.cn id=491 lang=golang
 * @lcpr version=21909
 *
 * [491] 递增子序列
 */

// @lc code=start
// 当题目中出现不能有重复的结果集，或者说出现了的重复元素会导致结果集出现重复的答案，此时就要考虑进行去重操作
package main

func findSubsequences(nums []int) [][]int {
	var path []int
	var result [][]int
	n := len(nums)
	var backTracking func(int)
	backTracking = func(startIndex int) {
		if len(path) >= 2 {
			result = append(result, append([]int(nil), path...))
		}
		if startIndex >= n {
			return
		}
		// 同一层的去重
		used := make(map[int]bool)
		for i := startIndex; i < n; i++ {
			if len(path) > 0 && nums[i] < path[len(path)-1] || used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return result
}

// if nums[i] >= last {
// 	path = append(path, nums[i])
// 	dfs(i+1, nums[i])
// 	path = path[:len(path)-1]
// }

// if nums[i] != last {
// 	dfs(i+1, last)
// }
// @lc code=end

/*
// @lcpr case=start
// [4,6,7,7]\n
// @lcpr case=end

// @lcpr case=start
// [4,4,3,2,1]\n
// @lcpr case=end

*/
