/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	cur := pre
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return pre.Next
}

// @lc code=end
