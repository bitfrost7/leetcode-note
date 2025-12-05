/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

package linkedlist

import . "leetcode-note/helpers"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

// @lc code=end
