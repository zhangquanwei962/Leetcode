/*
 * @lc app=leetcode.cn id=1669 lang=golang
 * @lcpr version=21909
 *
 * [1669] 合并两个链表
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

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	pre1 := list1
	for i := 0; i < a-1; i++ {
		pre1 = pre1.Next
	}
	pre2 := pre1
	for i := 0; i < b-a+2; i++ {
		pre2 = pre2.Next
	}
	pre1.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = pre2
	return list1
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,2,3,4,5]\n3\n4\n[1000000,1000001,1000002]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,3,4,5,6]\n2\n5\n[1000000,1000001,1000002,1000003,1000004]\n
// @lcpr case=end

*/
