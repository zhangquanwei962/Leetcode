/*
 * @lc app=leetcode.cn id=907 lang=golang
 * @lcpr version=21909
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
package main

func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	// 左边界 left[i] 为左侧严格小于 arr[i] 的最近元素位置（不存在时为 -1）
	left := make([]int, n)
	st := []int{-1} // 方便赋值 left
	for i, x := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= x {
			st = st[:len(st)-1] // 移除无用数据
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	// 右边界 right[i] 为右侧小于等于 arr[i] 的最近元素位置（不存在时为 n）
	right := make([]int, n)
	st = []int{n} // 方便赋值 right
	for i := n - 1; i >= 0; i-- {
		for len(st) > 1 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1] // 移除无用数据
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	for i, x := range arr {
		ans += x * (i - left[i]) * (right[i] - i) // 累加贡献
	}
	return ans % mod
}

/*
func sumSubarrayMins(arr []int) (ans int) {
    n := len(arr)
    left := make([]int, n)
    right := make([]int, n)
    for i := range right {
        right[i] = n
    }
    st := []int{-1} // 方便赋值 left
    for i, x := range arr {
        for len(st) > 1 && arr[st[len(st)-1]] >= x {
            right[st[len(st)-1]] = i // i 恰好是栈顶的右边界
            st = st[:len(st)-1]
        }
        left[i] = st[len(st)-1]
        st = append(st, i)
    }

    for i, x := range arr {
        ans += x * (i - left[i]) * (right[i] - i) // 累加贡献
    }
    return ans % (1e9 + 7)
}
*/

// @lc code=end

/*
// @lcpr case=start
// [3,1,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [11,81,94,43,3]\n
// @lcpr case=end

*/
