package dynamic_programming

import (
	"fmt"
	"math"
)

/*
https://leetcode.cn/problems/minimum-falling-path-sum/description
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。
在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1)
*/

// 由下往上递推即可
// 每个元素的最小下降路径等于本元素加上下一行最小的元素
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if j == 0 {
				matrix[i][j] = min(matrix[i+1][j], matrix[i+1][j+1]) + matrix[i][j]
			} else if j == n-1 {
				matrix[i][j] = min(matrix[i+1][j], matrix[i+1][j-1]) + matrix[i][j]
			} else {
				matrix[i][j] = min(min(matrix[i+1][j+1], matrix[i+1][j]), matrix[i+1][j-1]) + matrix[i][j]
			}
		}
	}
	res := math.MaxInt
	for _, v := range matrix[0] {
		res = min(res, v)
	}
	return res
}
func TestMinFallingPathSum() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(matrix))
}
