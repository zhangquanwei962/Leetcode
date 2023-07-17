/*
 * @lc app=leetcode.cn id=445 lang=golang
 * @lcpr version=21909
 *
 * [445] 两数相加 II
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stk1, stk2 []int
	for l1 != nil {
		stk1 = append(stk1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stk2 = append(stk2, l2.Val)
		l2 = l2.Next
	}

	carry := 0
	var head *ListNode

	for len(stk1) > 0 || len(stk2) > 0 || carry != 0 {
		n1, n2 := 0, 0
		if len(stk1) > 0 {
			n1 = stk1[len(stk1)-1]
			stk1 = stk1[:len(stk1)-1]
		}

		if len(stk2) > 0 {
			n2 = stk2[len(stk2)-1]
			stk2 = stk2[:len(stk2)-1]
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		curnode := &ListNode{Val: sum}
		curnode.Next = head // 更新当前节点的下一个节点
		head = curnode      // 更新头节点为当前节点
	}

	return head
}

// @lc code=end

/*
// @lcpr case=start
// [7,2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

*/
