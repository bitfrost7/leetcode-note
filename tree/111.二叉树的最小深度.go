/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	stack := []*TreeNode{root}
	depth := 0
	for len(stack) > 0 {
		depth++
		l := len(stack)
		for _, node := range stack {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[l:]
	}
	return depth
}

// @lc code=end
