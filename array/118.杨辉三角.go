package array

import "fmt"

// description: (https://leetcode-cn.com/problems/pascals-triangle/)
// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
// 示例 1:
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

// ----- Code Here -----
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	res := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		row := make([]int, 1, i+1)
		row[0] = 1
		for j := 1; j < i; j++ {
			row = append(row, res[i-1][j-1]+res[i-1][j])
		}
		row = append(row, 1)
		res = append(res, row)
	}
	return res
}

// ----- Begin Test -----
func TestGenerate() {
	fmt.Println("Input: numRows = 5")
	fmt.Println(generate(5))
}

// ----- Solution Here -----

// 解法：动态生成每一行
//
// 杨辉三角第 i 行（0-indexed）长度为 i+1：
// - 每行第一个和最后一个元素固定为 1
// - 中间位置 j 的值来自上一行：res[i-1][j-1] + res[i-1][j]
//
// 复杂度：
// - 时间 O(numRows^2)
// - 额外空间 O(numRows^2)（返回结果占用）
