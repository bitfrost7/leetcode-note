/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	node := root
	res := make([]int, 0)
	for node != nil || len(stack) > 0 {
		// 一直向左走，直到左子树为空
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// 弹出栈顶元素
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 访问栈顶元素
		res = append(res, node.Val)
		// 访问右子树
		node = node.Right
	}
	return res
}

// @lc code=end
