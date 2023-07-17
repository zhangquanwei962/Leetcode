/*
 * @lc app=leetcode.cn id=2178 lang=golang
 * @lcpr version=21909
 *
 * [2178] 拆分成最多数目的正偶数之和
 */

// @lc code=start
// 贪心，尽可能地多分，当finalSum<i时候，把剩余的加上去
package main

func maximumEvenSplit(finalSum int64) (ans []int64) {
	if finalSum&1 == 1 {
		return
	}

	for i := int64(2); i <= finalSum; i += 2 {
		ans = append(ans, int64(i))
		finalSum -= i
	}
	ans[len(ans)-1] += finalSum
	return
}

// @lc code=end

/*
// @lcpr case=start
// 12\n
// @lcpr case=end

// @lcpr case=start
// 7\n
// @lcpr case=end

// @lcpr case=start
// 28\n
// @lcpr case=end

*/
