/*
 * @lc app=leetcode.cn id=2223 lang=golang
 * @lcpr version=21909
 *
 * [2223] 构造字符串的总得分和
 */

// @lc code=start
// O(nlogn) O(n)
package main

const P uint64 = 131

func sumScores(s string) int64 {
	n := len(s)
	h := make([]uint64, n+1)
	p := make([]uint64, n+1)
	p[0] = 1
	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*P + uint64(s[i-1])
		p[i] = p[i-1] * P
	}

	get := func(l, r int) uint64 {
		return h[r] - h[l-1]*p[r-l+1]
	}
	ans := n
	for length := 1; length < n; length++ {
		if s[n-length] != s[0] {
			continue
		}

		left, right := -1, length+1
		for left+1 < right {
			mid := left + (right-left)/2
			if get(n-length+1, n-length+mid) == get(1, mid) {
				left = mid
			} else {
				right = mid
			}
		}
		ans += left
	}
	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// "babab"\n
// @lcpr case=end

// @lcpr case=start
// "azbazbzaz"\n
// @lcpr case=end

*/
