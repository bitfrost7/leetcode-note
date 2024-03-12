package linkedlist

// https://leetcode.cn/problems/swap-nodes-in-pairs/description/
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(head.Next.Next)
	next.Next = head
	return next
}

func TestSwapPairs() {
	list := []int{1, 2, 3, 4}
	head := NewLinkedList(list)
	PrintLinkedList(head, "Input")
	PrintLinkedList(swapPairs(head), "交换后链表")
}
