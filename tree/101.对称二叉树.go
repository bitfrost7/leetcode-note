/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
package tree

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	l, r := root.Left, root.Right
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return isSymmetricHelper(l, r)
}

func isSymmetricHelper(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return isSymmetricHelper(l.Left, r.Right) && isSymmetricHelper(l.Right, r.Left)
}

// @lc code=end
