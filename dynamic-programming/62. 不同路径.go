package dynamic_programming

import "fmt"

/*
https://leetcode.cn/problems/unique-paths/description
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

示例 1：
输入：m = 3, n = 7
输出：28

示例 2：
输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		dp[i][0] = 1
	}
	for j := range dp[0] {
		dp[0][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	printArr2(dp)
	return dp[n-1][m-1]
}
func uniquePaths1(m int, n int) int {
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
	m := 3
	n := 7
	fmt.Println("input: m=", m, " n=", n)
	fmt.Printf("从左上角开始，总共有 %d 条路径可以到达右下角。", uniquePaths(m, n))
}
