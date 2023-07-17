/*
 * @lc app=leetcode.cn id=1130 lang=golang
 * @lcpr version=21908
 *
 * [1130] 叶值的最小代价生成树
 */

// @lc code=start
package main

import "fmt"

func mctFromLeafValues(arr []int) (ans int) {
	stk := make([]int, len(arr))
	for _, x := range arr {
		for len(stk) > 0 && stk[len(stk)-1] <= x {
			if len(stk) == 1 || stk[len(stk)-2] > x { //为空或者栈顶大于x
				ans += stk[len(stk)-1] * x
			} else {
				ans += stk[len(stk)-2] * stk[len(stk)-1] //y与栈顶合并
			}
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, x)
	}
	fmt.Println(stk, ans)
	for len(stk) > 1 {
		ans += stk[len(stk)-2] * stk[len(stk)-1]
		stk = stk[:len(stk)-1]
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [6,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [4,11]\n
// @lcpr case=end

*/
