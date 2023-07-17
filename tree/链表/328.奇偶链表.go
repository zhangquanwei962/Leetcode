/*
 * @lc app=leetcode.cn id=328 lang=golang
 * @lcpr version=21909
 *
 * [328] 奇偶链表
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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	count := 1
	for ; head != nil; head = head.Next {
		if count&1 == 1 {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		count++
	}

	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

// func oddEvenList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	evenHead := head.Next
// 	odd := head
// 	even := evenHead

// 	for even != nil && even.Next != nil {
// 		odd.Next = even.Next
// 		odd = odd.Next
// 		even.Next = odd.Next
// 		even = even.Next
// 	}
// 	odd.Next = evenHead
// 	return head
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,3,5,6,4,7]\n
// @lcpr case=end

*/
