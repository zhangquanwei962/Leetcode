/*
 * @lc app=leetcode.cn id=1239 lang=golang
 * @lcpr version=21909
 *
 * [1239] 串联字符串的最大长度
 */

// @lc code=start
// 2^16暴力位运算枚举 回 溯
package main

import "math/bits"

func maxLength(arr []string) (ans int) {
	masks := []int{}
outer:
	for _, s := range arr {
		mask := 0
		for _, ch := range s {
			ch -= 'a'
			if mask>>ch&1 > 0 { // 若 mask 已有 ch，则说明 s 含有重复字母，无法构成可行解
				continue outer
			}
			mask |= 1 << ch // 将 ch 加入 mask 中
		}
		masks = append(masks, mask)
	}

	n := len(masks)
	var dfs func(int, int)
	dfs = func(pos, mask int) {
		// if pos == n {
		ans = max(ans, bits.OnesCount(uint(mask)))
		// 	return
		// }
		for i := pos; i < n; i++ {
			if mask&masks[i] != 0 {
				continue
			}
			mask |= masks[i]
			dfs(i+1, mask)
			mask ^= masks[i]
		}
		// // 选
		// if mask&masks[pos] == 0 { // mask 和 masks[pos] 无公共元素
		// 	dfs(pos+1, mask|masks[pos])
		// }
		// // 不选
		// dfs(pos+1, mask)
	}
	dfs(0, 0)
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
// ["un","iq","ue"]\n
// @lcpr case=end

// @lcpr case=start
// ["cha","r","act","ers"]\n
// @lcpr case=end

// @lcpr case=start
// ["abcdefghijklmnopqrstuvwxyz"]\n
// @lcpr case=end

*/
