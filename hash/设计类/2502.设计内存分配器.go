/*
 * @lc app=leetcode.cn id=2502 lang=golang
 * @lcpr version=21909
 *
 * [2502] 设计内存分配器
 */

// @lc code=start
package main

type Allocator []int

func Constructor(n int) Allocator {
	return make(Allocator, n)
}

func (a Allocator) Allocate(size int, mID int) int {
	cnt := 0
	for i, id := range a {
		if id > 0 {
			cnt = 0
		} else if cnt++; cnt == size {
			for j := i; j > i-size; j-- {
				a[j] = mID
			}
			return i - size + 1
		}
	}
	return -1

}

func (a Allocator) Free(mID int) (ans int) {
	for i, id := range a {
		if id == mID {
			ans++
			a[i] = 0
		}
	}
	return
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */
// @lc code=end
