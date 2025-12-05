package linkedlist

import . "leetcode-note/helpers"

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	p := head
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		p.Next = &ListNode{
			Val: sum % 10,
		}
		p = p.Next
	}
	if carry == 1 {
		p.Next = &ListNode{
			Val: 1,
		}
	}
	return head.Next
}

func TestAddTwoNumbers() {
	nums1 := []int{2, 4, 5, 9}
	nums2 := []int{5, 6, 4}
	l1 := NewLinkedList(nums1)
	PrintLinkedList(l1, "l1")
	l2 := NewLinkedList(nums2)
	PrintLinkedList(l2, "l2")
	PrintLinkedList(addTwoNumbers(l1, l2), "l1+l2")
}
