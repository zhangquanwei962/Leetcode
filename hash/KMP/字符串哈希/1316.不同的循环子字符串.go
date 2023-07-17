/*
 * @lc app=leetcode.cn id=1316 lang=golang
 * @lcpr version=21909
 *
 * [1316] 不同的循环子字符串
 */

// @lc code=start
package main

const P int64 = 131

func distinctEchoSubstrings(text string) (ans int) {
	n := len(text)
	h := make([]int64, n+1)
	p := make([]int64, n+1)
	p[0] = 1
	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*P + int64(text[i-1])
		p[i] = p[i-1] * P
	}

	get := func(l, r int) int {
		return int(h[r] - h[l-1]*p[r-l+1])
	}

	st := make(map[int]struct{})
	for len := 1; 2*len <= n; len++ { // 枚举长度
		for i := 1; i+2*len-1 <= n; i++ { // 枚举起点
			j := i + len - 1 // 前一半终点
			a := get(i, j)
			b := get(j+1, j+len) // 后终点
			if a == b {
				if _, ok := st[a]; !ok {
					st[a] = struct{}{}
				}
			}
		}
	}
	return len(st)
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcabc"\n
// @lcpr case=end

// @lcpr case=start
// "leetcodeleetcode"\n
// @lcpr case=end

*/
