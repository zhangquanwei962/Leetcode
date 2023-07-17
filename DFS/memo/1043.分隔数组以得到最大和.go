/*
 * @lc app=leetcode.cn id=1043 lang=golang
 * @lcpr version=21907
 *
 * [1043] 分隔数组以得到最大和
 */
/*
既然是在f的最左边插入一个状态，那么就只需要修改和
f 有关的下标，其余任何逻辑都无需修改。*/
// @lc code=start
/*
func maxSumAfterPartitioning(arr []int, k int) int {
    n := len(arr)
    f := make([]int, n+1)
    for i := range arr {
        for j, mx := i, 0; j > i-k && j >= 0; j-- {
            mx = max(mx, arr[j]) // 一边枚举 j，一边计算子数组的最大值
            f[i+1] = max(f[i+1], f[j]+(i-j+1)*mx)
        }
    }
    return f[n]
}

func max(a, b int) int { if a < b { return b }; return a }
*/
package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i < 0 {
			return
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}

		defer func() { *p = res }()

		for j, mx := i, 0; j > i-k && j >= 0; j-- {
			mx = max(mx, arr[j]) // 一边枚举 j，一边计算子数组的最大值
			res = max(res, dfs(j-1)+(i-j+1)*mx)
		}
		return
	}
	return dfs(n - 1)
}

// @lc code=end

/*
// @lcpr case=start
// [1,15,7,9,2,5,10]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,4,1,5,7,3,6,1,9,9,3]\n4\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/
