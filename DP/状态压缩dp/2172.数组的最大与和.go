/*
 * @lc app=leetcode.cn id=2172 lang=golang
 * @lcpr version=21909
 *
 * [2172] 数组的最大与和
 */

// @lc code=start
// 视为有2*num个篮子，每个篮子一个整数
package main

import "math/bits"

func maximumANDSum(nums []int, numSlots int) (ans int) {
	f := make([]int, 1<<(numSlots*2))
	for i, fi := range f {
		c := bits.OnesCount(uint(i))
		if c >= len(nums) {
			continue
		}
		for j := 0; j < numSlots*2; j++ {
			if i>>j&1 == 0 { // 枚举空篮子 j
				s := i | 1<<j // 添加元素
				f[s] = max(f[s], fi+(j/2+1)&nums[c])
				ans = max(ans, f[s])
			}
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,3,10,4,7,1]\n9\n
// @lcpr case=end

*/
