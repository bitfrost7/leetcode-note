package dynamic_programming

import (
	"fmt"
	//lint:ignore ST1001 方便使用公共工具
	. "leetcode-note/helpers"
)

// https://leetcode.cn/problems/unique-binary-search-trees-ii/description
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

var DP = [10][10][]*TreeNode{} // DP[i][j] 代表从第i到第j个节点互不相同的不同 二叉搜索树
func generateTrees(n int) []*TreeNode {
	trees(1, n)
	return DP[1][n]
}
func trees(start, end int) {
	if DP[start][end] != nil {
		return
	}
	if start > end {
		DP[start][end] = []*TreeNode{nil}
		return
	}
	if start == end {
		DP[start][end] = []*TreeNode{{Val: start}}
		return
	}
	for i := start; i <= end; i++ {
		trees(start, i-1)
		trees(i+1, end)
		for _, left := range DP[start][i-1] {
			for _, right := range DP[i+1][end] {
				DP[start][end] = append(DP[start][end], &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
}

func TestGenerateTrees() {
	n := 3
	fmt.Printf("Input: n=%d\n", n)
	fmt.Println("组成二叉搜索树:")
	PrintTrees(generateTrees(n))
}
