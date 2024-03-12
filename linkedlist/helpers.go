package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return head.Next
}

func PrintLinkedList(head *ListNode, name string) {
	if len(name) > 0 {
		fmt.Print(name, ":\t")
	}
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("nil")
}
