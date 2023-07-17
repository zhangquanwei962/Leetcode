/*
 * @lc app=leetcode.cn id=2434 lang=golang
 * @lcpr version=21909
 *
 * [2434] 使用机器人打印字典序最小的字符串
 */

// @lc code=start
// **找后面的最小值**
// 一般的做法是倒着遍历一遍，用一个数组记录最小值。
// 对于本题来说，由于值域比较小（26），也可以用计数的方式算出来，内存消耗要小一些。
// 单调栈是为了找到比自身左边小的最近的一个位置
// 贪心栈是为了找到包含自身的，先append
package main

func robotWithString(s string) string {
	ans := make([]byte, 0, len(s))
	cnt := [26]int{}
	for _, x := range s {
		cnt[x-'a']++
	}

	min := byte(0) // 剩余的最小字母
	stk := []byte{}
	for _, c := range s {
		cnt[c-'a']--
		for min < 25 && cnt[min] == 0 { // 找到第一个正数
			min++
		}
		stk = append(stk, byte(c))
		for len(stk) > 0 && stk[len(stk)-1]-'a' <= min {
			ans = append(ans, stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}
	}

	return string(ans)
}

// @lc code=end

/*
// @lcpr case=start
// "zza"\n
// @lcpr case=end

// @lcpr case=start
// "bac"\n
// @lcpr case=end

// @lcpr case=start
// "bdda"\n
// @lcpr case=end

*/
