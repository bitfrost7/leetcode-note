package array

import "fmt"

// description: (https://leetcode-cn.com/problems/pascals-triangle/)
// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
// 示例 1:
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
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

func TestGenerate() {
	fmt.Println("Input: numRows = 5")
	fmt.Println(generate(5))
}
