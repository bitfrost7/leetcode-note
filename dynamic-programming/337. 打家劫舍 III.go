package dynamic_programming

// https://leetcode.cn/problems/house-robber-iii/description
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
//除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
//给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func robIII(root *TreeNode) int {
	amount := robTree(root)
	return max(amount[0], amount[1])
}

func robTree(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}
	l, r := robTree(root.Left), robTree(root.Right)
	do := root.Val + l[0] + r[0]
	dont := max(l[0], l[1]) + max(r[0], r[1])
	return [2]int{dont, do}
}
