package array

import "fmt"

// description: (https://leetcode-cn.com/problems/pascals-triangle-ii/)
// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
// 示例 1:
// 输入: rowIndex = 3
// 输出: [1,3,3,1]

// ----- Code Here -----
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

// ----- Begin Test -----
func TestGetRow() {
	fmt.Println("Input: rowIndex = 3")
	fmt.Println(getRow(3))
}

// ----- Solution Here -----

// 解法：一维数组原地更新（倒序）
//
// 目标是返回第 rowIndex 行，使用一维数组 row 逐行迭代构建：
// - row[0] 固定为 1
// - 第 i 行更新时，为了不覆盖 row[j-1]（上一轮的值），j 必须从 i 倒序到 1：
//   row[j] = row[j] + row[j-1]
//
// 复杂度：
// - 时间 O(rowIndex^2)
// - 额外空间 O(rowIndex)
