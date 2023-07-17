/*
 * @lc app=leetcode.cn id=134 lang=golang
 * @lcpr version=21908
 *
 * [134] 加油站
 */

// @lc code=start
package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	curGas, totalGas, startStation := 0, 0, 0
	for i := 0; i < n; i++ {
		curGas += gas[i] - cost[i]
		totalGas += gas[i] - cost[i]

		if curGas < 0 {
			curGas = 0
			startStation = i + 1
		}
	}

	if totalGas < 0 {
		return -1
	}
	return startStation
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n[3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n[3,4,3]\n
// @lcpr case=end

*/
