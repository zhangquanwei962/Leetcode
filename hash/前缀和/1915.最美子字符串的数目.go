/*
 * @lc app=leetcode.cn id=1915 lang=golang
 * @lcpr version=21909
 *
 * [1915] 最美子字符串的数目
 */

// @lc code=start
package main

func wonderfulSubstrings(word string) (ans int64) {
	cnt := [1024]int{1}
	sum := 0
	for _, c := range word {
		sum ^= 1 << (c - 'a')
		ans += int64(cnt[sum])

		for i := 1; i < 1024; i <<= 1 {
			ans += int64(cnt[sum^i])
		}
		cnt[sum]++
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "aabb"\n
// @lcpr case=end

// @lcpr case=start
// "he"\n
// @lcpr case=end

*/
