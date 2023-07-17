/*
 * @lc app=leetcode.cn id=93 lang=golang
 * @lcpr version=21908
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
package main

func restoreIpAddresses(s string) (res []string) {
	n := len(s)
	if n < 4 || n > 12 {
		return res
	}

	var dfs func(string, int, int)
	dfs = func(cur string, idx int, count int) {
		if idx == n && count == 4 {
			res = append(res, cur[:len(cur)-1])
			return
		}
		// 剪枝，不够的时候返回
		if n-idx > (4-count)*3 || n-idx < 4-count {
			return
		}

		num := 0
		for i := idx; i < idx+3 && i < n; i++ {
			num = num*10 + int(s[i]-'0')
			if num <= 255 {
				cur += string(s[i])
				dfs(cur+".", i+1, count+1)
			}
			if num == 0 {
				break
			}
		}
	}
	dfs("", 0, 0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// "25525511135"\n
// @lcpr case=end

// @lcpr case=start
// "0000"\n
// @lcpr case=end

// @lcpr case=start
// "101023"\n
// @lcpr case=end

*/
