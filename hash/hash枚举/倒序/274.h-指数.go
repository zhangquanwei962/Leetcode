/*
 * @lc app=leetcode.cn id=274 lang=golang
 * @lcpr version=21908
 *
 * [274] H 指数
 */

// @lc code=start
// counter用来记录当前引用次数的论文有几篇
// O(n) O(n)
package main

func hIndex(citations []int) int {
	n := len(citations)
	counter := make([]int, n+1)

	for _, cation := range citations {
		if cation >= n {
			counter[n]++
		} else {
			counter[cation]++
		}
	}

	for i, tot := n, 0; i >= 0; i-- {
		tot += counter[i]
		if tot >= i {
			return i
		}
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// [3,0,6,1,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,1]\n
// @lcpr case=end

*/
