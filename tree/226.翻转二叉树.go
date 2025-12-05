/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package tree

//lint:ignore ST1001 方便使用公共工具
import . "leetcode-note/helpers"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// @lc code=end

func TestInvertTree() {
	root := NewTreeNode(4, NewTreeNode(2, NewTreeNode(1, nil, nil), NewTreeNode(3, nil, nil)), NewTreeNode(7, NewTreeNode(6, nil, nil), NewTreeNode(9, nil, nil)))
	PrintTree(invertTree(root))
}
