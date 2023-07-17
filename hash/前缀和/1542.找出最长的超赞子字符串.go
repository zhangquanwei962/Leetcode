/*
 * @lc app=leetcode.cn id=1542 lang=golang
 * @lcpr version=21909
 *
 * [1542] 找出最长的超赞子字符串
 */

// @lc code=start
package main

func longestAwesome(s string) (ans int) {
	cnt := map[int]int{0: -1}
	sum := 0

	for i, ch := range s {
		sum ^= 1 << (ch - '0')

		if _, ok := cnt[sum]; ok {
			ans = max(ans, i-cnt[sum])
		} else {
			cnt[sum] = i
		}

		for k := 0; k < 10; k++ {
			newValue := sum ^ (1 << k)
			if _, ok := cnt[newValue]; ok {
				ans = max(ans, i-cnt[newValue])
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// "3242415"\n
// @lcpr case=end

// @lcpr case=start
// "12345678"\n
// @lcpr case=end

// @lcpr case=start
// "213123"\n
// @lcpr case=end

// @lcpr case=start
// "00"\n
// @lcpr case=end

*/
