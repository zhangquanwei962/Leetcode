/*
 * @lc app=leetcode.cn id=1031 lang=golang
 * @lcpr version=21908
 *
 * [1031] 两个非重叠子数组的最大和
 */

// @lc code=start
// 同向双指针 s[0]=0 s[1]-s[0]=nums[0]
// [left,right)是s[right]-s[left],长度为right-left
package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) (ans int) {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x // 前缀和
	}

	f := func(firstLen, secondLen int) {
		maxSumA := 0
		for i := firstLen + secondLen; i <= n; i++ {
			maxSumA = max(maxSumA, s[i-secondLen]-s[i-secondLen-firstLen])
			ans = max(ans, maxSumA+s[i]-s[i-secondLen])
		}
	}

	f(firstLen, secondLen) // 左 a 右 b
	f(secondLen, firstLen) // 左 b 右 a
	return
}

// @lc code=end

/*
// @lcpr case=start
// [0,6,5,2,2,5,1,9,4]\n1\n2\n
// @lcpr case=end

// @lcpr case=start
// [3,8,1,3,2,1,8,9,0]\n3\n2\n
// @lcpr case=end

// @lcpr case=start
// [2,1,5,6,0,9,5,0,3,8]\n4\n3\n
// @lcpr case=end

*/
