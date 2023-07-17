/*
 * @lc app=leetcode.cn id=1049 lang=golang
 * @lcpr version=21908
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
package main

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, x := range stones {
		sum += x
	}

	target := sum / 2

	f := make([]bool, target+1)
	f[0] = true

	for _, x := range stones {
		for j := target; j >= x; j-- {
			f[j] = f[j] || f[j-x]
		}
	}

	for i := target; i >= 0; i-- {
		if f[i] == true {
			target = i
			break
		}
	}

	return sum - 2*target

}

// @lc code=end

/*
// @lcpr case=start
// [2,7,4,1,8,1]\n
// @lcpr case=end

// @lcpr case=start
// [31,26,33,21,40]\n
// @lcpr case=end

*/
