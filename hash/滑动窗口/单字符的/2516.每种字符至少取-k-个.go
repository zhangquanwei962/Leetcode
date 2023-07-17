/*
 * @lc app=leetcode.cn id=2516 lang=golang
 * @lcpr version=21908
 *
 * [2516] 每种字符至少取 K 个
 */

// @lc code=start
package main

func takeCharacters(s string, k int) int {
	count := make(map[byte]int)
	count['a'], count['b'], count['c'] = 0, 0, 0
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	count['a'] -= k
	count['b'] -= k
	count['c'] -= k
	for _, v := range count {
		if v < 0 {
			return -1
		}
	}
	ans, left := 0, 0
	cnt := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		cnt[s[right]]++
		for cnt[s[right]] > count[s[right]] {
			cnt[s[left]]--
			left++
		}
		curLen := right - left + 1
		if curLen > ans {
			ans = curLen
		}
	}
	return len(s) - ans
}

// func takeCharacters(s string, k int) int {
// 	n := len(s)
// 	c, j := [3]int{}, n
// 	for c[0] < k || c[1] < k || c[2] < k {
// 		if j == 0 {
// 			return -1 // 所有字母都取也无法满足要求
// 		}
// 		j--
// 		c[s[j]-'a']++
// 	}

// 	ans := n - j //左侧没有取字符

// 	for i := 0; i < n && j < n; i++ { //双指针
// 		c[s[i]-'a']++
// 		for j < n && c[s[j]-'a'] > k { //维护j的最大下标
// 			c[s[j]-'a']--
// 			j++
// 		}
// 		ans = min(ans, i+1+n-j)
// 	}
// 	return ans
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

/*
// @lcpr case=start
// "aabaaaacaabc"\n2\n
// @lcpr case=end

// @lcpr case=start
// "a"\n1\n
// @lcpr case=end

*/
