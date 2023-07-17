/*
 * @lc app=leetcode.cn id=470 lang=golang
 * @lcpr version=21909
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
// 拒绝采样 变进制数
package main

func rand10() int {
	var first, second int
	for {
		first = rand7()
		if first <= 2 {
			break
		}
	}

	for {
		second = rand7()
		if second <= 5 {
			break
		}
	}

	return (first-1)*5 + second
}

// func rand10() int {
// 	var first, second int
// 	for {
// 		first = rand7()
// 		if first <= 6 {
// 			break
// 		}
// 	}
// 	for {
// 		second = rand7()
// 		if second <= 5 {
// 			break
// 		}
// 	}
// 	if first&1 == 1 {
// 		return second
// 	} else {
// 		return 5 + second
// 	}
// }

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/
