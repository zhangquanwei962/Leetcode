/*
 * @lc app=leetcode.cn id=2451 lang=golang
 * @lcpr version=21908
 *
 * [2451] 差值数组不同的字符串
 */

// @lc code=start
package main

func oddString(words []string) string {
	d := map[string][]string{}

	for _, s := range words {
		m := len(s)
		cs := make([]byte, m-1)
		for i := 0; i < m-1; i++ {
			cs[i] = s[i+1] - s[i]
		}
		t := string(cs)
		d[t] = append(d[t], s)
	}

	for _, ss := range d {
		if len(ss) == 1 {
			return ss[0]
		}
	}
	return ""
}

// @lc code=end

/*
// @lcpr case=start
// ["adc","wzy","abc"]\n
// @lcpr case=end

// @lcpr case=start
// ["aaa","bob","ccc","ddd"]\n
// @lcpr case=end

*/
