/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 */
package tree

//lint:ignore ST1001 方便使用公共工具
import . "leetcode-note/helpers"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q { //p和q都是nil
		return true
	}
	if p == nil || q == nil { //p和q有一个是nil
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}

// @lc code=end
