package dynamic_programming

import (
	"fmt"
	"math"
)

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
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(matrix))
}
