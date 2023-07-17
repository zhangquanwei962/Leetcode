/*
 * @lc app=leetcode.cn id=2526 lang=golang
 * @lcpr version=21909
 *
 * [2526] 找到数据流中的连续整数
 */

// @lc code=start
package main

type DataStream struct{ value, k, cnt int }

func Constructor(value, k int) DataStream {
	return DataStream{value, k, 0}
}

func (d *DataStream) Consec(num int) bool {
	if num == d.value {
		d.cnt++
	} else {
		d.cnt = 0
	}
	return d.cnt >= d.k
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
// @lc code=end

/*
// @lcpr case=start
// ["DataStream", "consec", "consec", "consec", "consec"][[4, 3], [4], [4], [4], [3]]\n
// @lcpr case=end

*/
