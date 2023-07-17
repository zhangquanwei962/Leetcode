/*
 * @lc app=leetcode.cn id=78 lang=golang
 * @lcpr version=21909
 *
 * [78] 子集
 */

// @lc code=start
package main

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 不选
		dfs(i + 1)

		// 选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return
}

// class Solution: 先append
//     def subsets(self, nums: List[int]) -> List[List[int]]:
//         ans = []
//         path = []
//         n = len(nums)

//         def dfs(i: int) -> None:
//             ans.append(path.copy())
//             if i == n:
//                 return
//             for j in range(i, n):
//                 path.append(nums[j])
//                 dfs(j + 1)
//                 path.pop()
//         dfs(0)
//         return ans
// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
