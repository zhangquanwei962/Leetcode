/*
 * @lc app=leetcode.cn id=1547 lang=golang
 * @lcpr version=21909
 *
 * [1547] 切棍子的最小成本
 */

// @lc code=start
package main

package main

import "math"

func minCost(n int, cuts []int) int {
    sort.Ints(cuts)
    cuts = append([]int{0}, append(cuts, n)...)
    memo := make(map[[2]int]int)
    var dfs func(int, int) int
    dfs = func(i int, j int) int {
        if i + 1 == j {
            return 0
        }
        if val, ok := memo[[2]int{i, j}]; ok {
            return val
        }
        ans := math.MaxInt32
        for k := i + 1; k < j; k++ {
            left := dfs(i, k)
            right := dfs(k, j)
            temp := cuts[j] - cuts[i] + left + right
            if temp < ans {
                ans = temp
            }
        }
        memo[[2]int{i, j}] = ans
        return ans
    }
    return dfs(0, len(cuts) - 1)
}


// func minCost(n int, cuts []int) int {
// 	sort.Sort(sort.IntSlice(cuts))
// 	cuts = append([]int{0}, append(cuts, n)...)
// 	memo := make([][]int, n+2)
// 	for i := range memo {
// 		memo[i] = make([]int, n+2)
// 		for j := range memo[i] {
// 			memo[i][j] = -1
// 		}
// 	}
// 	var dfs func(int, int) int
// 	dfs = func(i int, j int) int {
// 		if i+1 == j {
// 			return 0
// 		}

// 		p := &memo[i][j]
// 		if *p != -1 {
// 			return *p
// 		}

// 		ans := math.MaxInt32
// 		for k := i + 1; k < j; k++ {
// 			left := dfs(i, k)
// 			right := dfs(k, j)
// 			temp := cuts[j] - cuts[i] + left + right
// 			if temp < ans {
// 				ans = temp
// 			}
// 		}
// 		defer func() { *p = ans }()
// 		return ans
// 	}
// 	return dfs(0, len(cuts)-1)
// }

// @lc code=end

/*
// @lcpr case=start
// 7\n[1,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// 9\n[5,6,1,4,2]\n
// @lcpr case=end

*/
