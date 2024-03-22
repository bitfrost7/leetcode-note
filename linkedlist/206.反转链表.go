package linkedlist

// 递归写法
func reverseList(head *ListNode) *ListNode {
	//如果当前节点位nil或者为尾节点 则直接返回该节点
	// 例如 1-> nil 无需翻转
	if head == nil || head.Next == nil {
		return head
	}
	// 现在考虑当前节点之后的节点都完成翻转
	// 例如 当前节点为2 ： 1 -> 2 <-3 <- 4 <- 5
	// 只需要将当前完成翻转的尾节点指向自己，并且设置自己指向nil即可
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 迭代写法
func reverseList1(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func TestReverseList() {
	nums := []int{1, 2, 3, 4, 5}
	list := NewLinkedList(nums)
	PrintLinkedList(list, "反转前")
	PrintLinkedList(reverseList1(list), "反转后")
}
