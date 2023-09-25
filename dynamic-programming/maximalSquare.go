package dynamic_programming

import "fmt"

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	var res int
	if m == 1 && n == 1 {
		if matrix[0][0] == '1' {
			return 1
		}
		return 0
	}
	arr := make([]int, n)
	for i := range arr {
		arr[i] = int(matrix[0][i] - '0')
		res = max(res, arr[i])
	}
	for i := 1; i < m; i++ {
		diag := arr[0]
		arr[0] = int(matrix[i][0] - '0')
		res = max(res, arr[0])
		for j := 1; j < n; j++ {
			tmp := arr[j]
			if matrix[i][j] == '1' {
				if arr[j] == 0 || arr[j-1] == 0 || diag == 0 {
					arr[j] = 1
				} else {
					arr[j] = min(min(arr[j], arr[j-1]), diag) + 1
				}
				res = max(res, arr[j])
			} else {
				arr[j] = 0
			}
			diag, tmp = tmp, diag
			res = max(res, arr[j])
		}
	}
	return res * res
}

func TestMaximalSquare() {
	matrix := [][]byte{
		{'1', '0'},
		{'0', '0'},
	}
	//matrix := [][]byte{
	//	{'1', '0', '1', '0', '0'},
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '1', '1', '1', '1'},
	//	{'1', '0', '0', '1', '0'},
	//}
	fmt.Println(maximalSquare(matrix))
}
