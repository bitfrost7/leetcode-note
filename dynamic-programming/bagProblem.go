package dynamic_programming

func bagProblem(w, v []int, W int) int {
	n := len(w)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {
			if j < w[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j-w[i-1]]+v[i-1], dp[i-1][j])
			}
		}
	}
	return dp[n][W]
}
