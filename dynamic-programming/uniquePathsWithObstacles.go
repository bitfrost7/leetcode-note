package dynamic_programming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	arr := make([]int, n)
	if obstacleGrid[0][0] != 1 {
		arr[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				arr[j] = 0
				continue
			}
			if j > 0 && obstacleGrid[i][j-1] == 0 {
				arr[j] += arr[j-1]
			}
		}
	}
	return arr[n-1]
}
