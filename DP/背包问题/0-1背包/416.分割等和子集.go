/*
 * @lc app=leetcode.cn id=416 lang=golang
 * @lcpr version=21907
 *
 * [416] 分割等和子集
 */

// @lc code=start
// 背包是顺序无关的
package main

func canPartition(nums []int) bool {
	s := 0
	for _, x := range nums {
		s += x
	}

	if s < 0 || s&1 == 1 {
		return false
	}

	s /= 2

	f := make([]bool, s+1)
	f[0] = true
	// 因为是j>=x时候背包才能更新
	for _, x := range nums {
		for j := s; j >= x; j-- {
			f[j] = f[j] || f[j-x]
		}
	}

	return f[s]
}

// func canPartition(nums []int) bool {
// 	s, n := 0, len(nums)
// 	for _, x := range nums {
// 		s += x
// 	}

// 	if s < 0 || s&1 == 1 {
// 		return false
// 	}

// 	s /= 2

// 	f := make([][]bool, n+1)
// 	for i := range f {
// 		f[i] = make([]bool, s+1)
// 	}

// 	f[0][0] = true
// 	for i, x := range nums {
// 		for c := 0; c <= s; c++ {
// 			if c < x {
// 				f[i+1][c] = f[i][c]
// 			} else {
// 				f[i+1][c] = f[i][c] || f[i][c-x]
// 			}
// 		}
// 	}
// 	return f[n][s]
// }

// @lc code=end

/*
// @lcpr case=start
// [1,5,11,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n
// @lcpr case=end

*/
