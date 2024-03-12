/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */
package linkedlist

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

// @lc code=end
