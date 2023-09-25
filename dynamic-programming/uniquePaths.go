package dynamic_programming

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

func TestUniquePaths() {
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 3))
}
