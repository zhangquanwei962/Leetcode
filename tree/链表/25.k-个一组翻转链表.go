/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=21909
 *
 * [25] K 个一组翻转链表
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++ // 统计节点个数
	}

	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre, cur *ListNode = nil, p0.Next

	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre // 每次循环只修改一个 Next，方便大家理解
			pre = cur
			cur = nxt
		}
		// 见视频
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt

	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/
