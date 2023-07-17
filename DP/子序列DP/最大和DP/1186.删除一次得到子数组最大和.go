/*
 * @lc app=leetcode.cn id=1186 lang=golang
 * @lcpr version=21909
 *
 * [1186] 删除一次得到子数组最大和
 */

// @lc code=start
// 和子数组最大和之间差距就是删除一次
// 子数组最大和还是要记录
// 要么不删除，维持之前的状态，要么删除掉一次，[0,..i-1]为
// 不删除掉的

// 要么是加上当前元素，也就是维持之前删除某个元素的情形，即g[i-1]+arr[i]
// 要么是删除当前这个元素，那么区间[0, i-1]就是不删除元素的情况，即f(i-1)+0（注意是f不是g！！）
package main

func maximumSum(arr []int) int {
	pre, maxn := 0, arr[0]
	res := arr[0]
	for _, x := range arr[1:] {
		pre = max(pre+x, maxn) // pre是维持之前状态
		// maxn是一直加的，少了x，删除x
		maxn = max(maxn+x, x)
		res = max(res, max(pre, maxn))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func maximumSum(arr []int) int {
//     ans := math.MinInt / 2
//     f0, f1 := ans, ans
//     for _, x := range arr {
//         f1 = max(f1+x, f0)
//         f0 = max(f0, 0) + x
//         ans = max(ans, max(f0, f1))
//     }
//     return ans
// }

// func max(a, b int) int { if b > a { return b }; return a }

// package main

// import "math"

// func maximumSum(arr []int) int {
// 	memo := make([][2]int, len(arr))
// 	for i := range memo {
// 		memo[i] = [2]int{math.MinInt, math.MinInt}
// 	}

// 	var dfs func(int, int) int
// 	dfs = func(i, j int) (res int) {
// 		if i < 0 {
// 			return math.MinInt / 2
// 		}
// 		p := &memo[i][j]
// 		if *p != math.MinInt { // 之前计算过
// 			return *p
// 		}
// 		defer func() { *p = res }() // 记忆化
// 		if j == 0 {
// 			return max(dfs(i-1, 0), 0) + arr[i]
// 		}

// 		return max(dfs(i-1, 1)+arr[i], dfs(i-1, 0))
// 	}
// 	ans := math.MinInt
// 	for i := range arr {
// 		ans = max(ans, max(dfs(i, 0), dfs(i, 1)))
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if b > a {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

/*
// @lcpr case=start
// [1,-2,0,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,-2,-2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,-1,-1,-1]\n
// @lcpr case=end

*/
