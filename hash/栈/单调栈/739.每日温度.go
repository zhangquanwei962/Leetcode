/*
 * @lc app=leetcode.cn id=739 lang=golang
 * @lcpr version=21909
 *
 * [739] 每日温度
 */

// @lc code=start
package main

func dailyTemperatures(temperatures []int) []int {
	var stk []int
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stk) > 0 && temperatures[i] >= temperatures[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			ans[i] = stk[len(stk)-1] - i
		} else {
			ans[i] = 0
		}
		stk = append(stk, i)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [73,74,75,71,69,72,76,73]\n
// @lcpr case=end

// @lcpr case=start
// [30,40,50,60]\n
// @lcpr case=end

// @lcpr case=start
// [30,60,90]\n
// @lcpr case=end

*/
