package dynamic_programming

import (
	"leetcode-note/helpers"
	. "leetcode-note/helpers"
	"math"
)

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description
// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
func maxPathSum(root *helpers.TreeNode) int {
	maxSum := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		pathSum := root.Val + left + right
		maxSum = max(maxSum, pathSum)
		return root.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}
