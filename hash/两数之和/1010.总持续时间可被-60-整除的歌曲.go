/*
 * @lc app=leetcode.cn id=1010 lang=golang
 * @lcpr version=21907
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lc code=start
// 有两个变量的题目，通常可以枚举一个
// 变量，把他当作常量，从而变成只有一个
// 变量的问题
package main

//(a+b)%c = (a %c + b % c) %c -> (60 - b % c) % c = a
func numPairsDivisibleBy60(time []int) (ans int) {
	cnt := [60]int{}
	for _, t := range time {
		// 先查询 cnt，再更新 cnt，因为题目要求 i<j
		// 如果先更新，再查询，就把 i=j 的情况也考虑进去了
		ans += cnt[(60-t%60)%60]
		cnt[t%60]++
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [30,20,150,100,40]\n
// @lcpr case=end

// @lcpr case=start
// [60,60,60]\n
// @lcpr case=end

*/
