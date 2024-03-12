package recursion

// 给你一个链表的头节点 head 。
//
//移除每个右侧有一个更大数值的节点。
//
//返回修改后链表的头节点 head 。

func removeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := removeNodes(head.Next)
	if next.Val > head.Val {
		return next
	}
	head.Next = next
	return head
}

func TestRemoveNodes() {
	head := NewLinkedList([]int{5, 2, 13, 3, 8})
	PrintLinkedList(head, "input")
	PrintLinkedList(removeNodes(head), "移除右侧有更大数值的节点后")
}
