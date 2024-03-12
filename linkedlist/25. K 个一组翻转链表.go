package linkedlist

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	var prev, cur, next *ListNode
	cur = head
	for i := 0; i < k; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	head.Next = reverseKGroup(cur, k)
	return prev
}

func TestReverse() {
	head := NewLinkedList([]int{1, 2, 3, 4, 5})
	PrintLinkedList(head, "反转前")
	PrintLinkedList(reverseKGroup(head, 2), "反转后")
}
