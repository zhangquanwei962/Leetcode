/*
 * @lc app=leetcode.cn id=24 lang=golang
 * @lcpr version=21909
 *
 * [24] 两两交换链表中的节点
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// func swapPairs(head *ListNode) *ListNode {
// 	dummy := &ListNode{Next: head}
// 	p0 := dummy
// 	var cur *ListNode = p0.Next

// 	for cur != nil && cur.Next != nil {
// 		nxt := cur.Next
// 		tmp := cur.Val
// 		cur.Val = nxt.Val
// 		nxt.Val = tmp
// 		cur = nxt.Next
// 	}
// 	return dummy.Next
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
