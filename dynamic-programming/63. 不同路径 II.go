package dynamic_programming

import (
	"fmt"

	. "leetcode-note/helpers"
)

/*
https://leetcode.cn/problems/unique-paths-ii/description
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右

输入：obstacleGrid = [[0,1],[0,0]]
输出：1
*/

// 动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 设dp[i][j] 为start到该点的路径和
	dp := make([][]int, m)
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}
	// 递推
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			// 边界条件
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
			}
			if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j]
			}
			// 状态转移
			if i != 0 && j != 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 空间优化到1维
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
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

func TestUniquePathsWithObstacles() {
	//obstacleGrid := [][]int{
	//	{0, 0, 0},
	//	{0, 1, 0},
	//	{0, 0, 0},
	//}
	obstacleGrid := [][]int{
		{1},
		{0},
	}
	fmt.Println("m x n 网格：")
	PrintArr2(obstacleGrid)
	fmt.Printf("从左上角到右下角一共有 %d 条不同的路径", uniquePathsWithObstacles(obstacleGrid))
}
