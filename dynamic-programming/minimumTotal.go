package dynamic_programming

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	arr := make([]int, len(triangle[n-1]))
	for i := range arr {
		arr[i] = triangle[n-1][i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			arr[j] = min(arr[j], arr[j+1]) + triangle[i][j]
		}
	}
	return arr[0]
}

func TestMinimumTotal() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
