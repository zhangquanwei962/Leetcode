/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=21909
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	var pre, cur *ListNode = nil, p0.Next
	for i := 0; i < right-left+1; i++ {
		nxt := cur.Next
		cur.Next = pre // 每次循环只修改一个 Next，方便大家理解
		pre = cur
		cur = nxt
	}

	// 见视频
	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/
