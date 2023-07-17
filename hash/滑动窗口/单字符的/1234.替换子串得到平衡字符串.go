/*
 * @lc app=leetcode.cn id=1234 lang=golang
 * @lcpr version=21909
 *
 * [1234] 替换子串得到平衡字符串
 */

// @lc code=start
package main

func balancedString(s string) int {
	cnt, m := ['X']int{}, len(s)/4 // 也可以用哈希表，不过数组更快一些
	for _, c := range s {
		cnt[c]++
	}
	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return 0 // 已经符合要求啦
	}
	ans, left := len(s), 0
	for right, c := range s { // 枚举子串右端点
		cnt[c]--
		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
			ans = min(ans, right-left+1)
			cnt[s[left]]++
			left++ // 缩小子串
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "QWER"\n
// @lcpr case=end

// @lcpr case=start
// "QQWE"\n
// @lcpr case=end

// @lcpr case=start
// "QQQW"\n
// @lcpr case=end

// @lcpr case=start
// "QQQQ"\n
// @lcpr case=end

*/
