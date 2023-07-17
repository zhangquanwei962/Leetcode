/*
 * @lc app=leetcode.cn id=6910 lang=golang
 * @lcpr version=21909
 *
 * [6910] 将数组划分成若干好子数组的方式
 */

// @lc code=start
// 乘法原理，每2个1之间划一条线，x个0可以划x+1条线
package main

func numberOfGoodSubarraySplits(nums []int) int {
	const mod int = 1e9 + 7
	ans, pre := 1, -1
	for i, x := range nums {
		if x > 0 {
			if pre >= 0 {
				ans = ans * (i - pre) % mod
			}
			pre = i
		}
	}
	if pre < 0 { // 整个数组都是 0，没有好子数组
		return 0
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0]\n
// @lcpr case=end

*/
