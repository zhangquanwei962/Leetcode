/*
 * @lc app=leetcode.cn id=202 lang=golang
 * @lcpr version=21909
 *
 * [202] 快乐数
 */

// @lc code=start
// O(logn) O(1)
// 快指针每次挑两个不用担心跳过1的点，因为1的点的下一个点还是1。快指针在慢指针后一个位置出发，这是因为只要求可追击，不要求环形列表的起始点。
package main

func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}

	return fast == 1

}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// 19\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/
