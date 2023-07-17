/*
 * @lc app=leetcode.cn id=1823 lang=golang
 * @lcpr version=21909
 *
 * [1823] 找出游戏的获胜者
 */

// @lc code=start
// 队列模拟 将队首元素取出并将该元素在队尾处重新加入队列
// 2.约瑟夫环
package main

// func findTheWinner(n int, k int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	ans := (findTheWinner(n-1, k) + k) % n
// 	if ans == 0 {
// 		return n
// 	}
// 	return ans
// }

func findTheWinner(n int, k int) (ans int) {
	for i := 2; i <= n; i++ {
		ans = (ans + k) % i
	}
	ans++ // 0~n-1映射到1~n
	return
}

// func findTheWinner(n int, k int) int {
// 	q := make([]int, n)
// 	for i := range q {
// 		q[i] = i + 1
// 	}
// 	for len(q) > 1 {
// 		for i := 1; i < k; i++ {
// 			q = append(q, q[0])[1:]
// 		}
// 		q = q[1:]
// 	}
// 	return q[0]
// }

// @lc code=end

/*
// @lcpr case=start
// 5\n2\n
// @lcpr case=end

// @lcpr case=start
// 6\n5\n
// @lcpr case=end

*/
