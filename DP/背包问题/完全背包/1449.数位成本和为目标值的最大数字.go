/*
 * @lc app=leetcode.cn id=1449 lang=golang
 * @lcpr version=21908
 *
 * [1449] 数位成本和为目标值的最大数字
 */

// @lc code=start\
package main

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func largestNumber(cost []int, target int) string {
	f := make([]int, target+1)

	for i := range f {
		f[i] = math.MinInt
	}
	f[0] = 0

	for _, x := range cost {
		for j := x; j <= target; j++ {
			f[j] = max(f[j], f[j-x]+1)
		}
	}

	if f[target] < 0 {
		return "0"
	}

	ans := make([]byte, 0, f[target])

	for i, j := 8, target; i >= 0; i-- {
		for c := cost[i]; j >= c && f[j] == f[j-c]+1; j -= c {
			ans = append(ans, byte('1'+i))
		}
	}
	return string(ans)

}

// @lc code=end

/*
// @lcpr case=start
// [4,3,2,5,6,7,2,5,5]\n9\n
// @lcpr case=end

// @lcpr case=start
// [7,6,5,5,5,6,8,7,8]\n12\n
// @lcpr case=end

// @lcpr case=start
// [2,4,6,2,4,6,4,4,4]\n5\n
// @lcpr case=end

// @lcpr case=start
// [6,10,15,40,40,40,40,40,40]\n47\n
// @lcpr case=end

*/
