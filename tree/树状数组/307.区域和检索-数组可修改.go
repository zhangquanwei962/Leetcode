/*
 * @lc app=leetcode.cn id=307 lang=golang
 * @lcpr version=21909
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
// 从1开始下标
package main

type NumArray struct {
	nums, tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	na := NumArray{nums, tree}
	for i, num := range nums {
		na.add(i+1, num)
	}
	return na
}

func (na *NumArray) add(index, val int) {
	for ; index < len(na.tree); index += index & -index {
		na.tree[index] += val
	}
}

func (na *NumArray) prefixSum(index int) (sum int) {
	for ; index > 0; index &= index - 1 {
		sum += na.tree[index]
	}
	return
}

func (na *NumArray) Update(index, val int) {
	na.add(index+1, val-na.nums[index])
	na.nums[index] = val
}

func (na *NumArray) SumRange(left, right int) int {
	return na.prefixSum(right+1) - na.prefixSum(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end

/*
// @lcpr case=start
// ["NumArray", "sumRange", "update", "sumRange"][[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]\n
// @lcpr case=end

*/
