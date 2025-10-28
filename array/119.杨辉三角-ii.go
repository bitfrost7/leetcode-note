package array

import "fmt"

// description: (https://leetcode-cn.com/problems/pascals-triangle-ii/)
// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
// 示例 1:
// 输入: rowIndex = 3
// 输出: [1,3,3,1]
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
	}
	return row
}

// 题解：
// 和杨辉三角相比主要是原地生成， 这样节省空间，核心点在于 原地生成下一组排列时候 需要倒序生成。

func TestGetRow() {
	fmt.Println("Input: rowIndex = 3")
	fmt.Println(getRow(3))
}
