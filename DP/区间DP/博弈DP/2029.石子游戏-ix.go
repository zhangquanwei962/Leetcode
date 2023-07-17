/*
 * @lc app=leetcode.cn id=2029 lang=golang
 * @lcpr version=21909
 *
 * [2029] 石子游戏 IX
 */

// @lc code=start
// 只关注能否被3整除，可以分组 0 1 2
// 枚举1 就是序列1121212....
package main

func check(c [3]int) bool {
	if c[1] == 0 { // 按照规矩，失败了
		return false
	}
	c[1]--                               // 开头为 1
	turn := 1 + min(c[1], c[2])*2 + c[0] // 计算回合数
	if c[1] > c[2] {                     // 序列末尾可以再加个 1
		turn++
		c[1]--
	}
	return turn&1 == 1 && c[1] != c[2] // 回合数为奇数，且还有剩余石子
}

func stoneGameIX(stones []int) bool {
	c := [3]int{}
	for _, v := range stones {
		c[v%3]++ //计数
	}
	return check(c) || check([3]int{c[0], c[2], c[1]}) // 枚举第一回合移除的是 1 还是 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [2,1]\n
// @lcpr case=end

// @lcpr case=start
// [2]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,2,4,3]\n
// @lcpr case=end

*/
