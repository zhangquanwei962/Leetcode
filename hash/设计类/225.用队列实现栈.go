/*
 * @lc app=leetcode.cn id=225 lang=golang
 * @lcpr version=21909
 *
 * [225] 用队列实现栈
 */

// @lc code=start
package main

type MyStack struct {
	q1, q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x)
	for len(this.q1) > 0 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}

	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
	v := this.q1[0]
	this.q1 = this.q1[1:]
	return v
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

/*
// @lcpr case=start
// ["MyStack", "push", "push", "top", "pop", "empty"][[], [1], [2], [], [], []]\n
// @lcpr case=end

*/
