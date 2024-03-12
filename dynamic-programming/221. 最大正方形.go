package dynamic_programming

import "fmt"

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
// 示例 1：
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// 输出：4
// 示例 2：
// 输入：matrix = [["0","1"],["1","0"]]
// 输出：1
// 示例 3：
// 输入：matrix = [["0"]]
// 输出：0
// f(i,j)表示以(i,j)为右下角的正方形的最大边长.
// 如果 (i,j)为“0”，以(i,j)为右下角不可能构成全为“1”的正方形f(i,j)=0
// 如果(i,j)为“1”，至少可以获得边长为1的正方形，还能不能变大只能向左向上扩展边长
// 这个时候需要看正上，左边和左上三个点，因为扩展定会将这三个相邻点包含进来，如果三个点中最小值为0，那么扩展后肯定不行
// 如果最小值为1，那么三个点都为1，定能扩展成边长为2的正方形，同理能扩展到最大的是 min(左，上，左上) + 1。

// 动态规划
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	maxSide := 0
	if m == 1 && n == 1 {
		if matrix[0][0] == '1' {
			return 1
		}
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSide = 1
		}
	}
	for j := range dp[0] {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxSide = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}

// 不需要处理边界
func maximalSquare1(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	maxSide := 0
	if m == 1 && n == 1 {
		if matrix[0][0] == '1' {
			return 1
		}
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}

// 空间优化
// 保存对角线的值
func maximalSquare2(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	var maxSide, diag int
	if m == 1 && n == 1 {
		if matrix[0][0] == '1' {
			return 1
		}
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i < m+1; i++ {
		diag = dp[0]
		for j := 1; j < n+1; j++ {
			tmp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = 1 + min(diag, min(dp[j], dp[j-1]))
			} else {
				dp[j] = 0
			}
			if dp[j] > maxSide {
				maxSide = dp[j]
			}
			diag = tmp
		}
	}
	return maxSide * maxSide
}

func TestMaximalSquare() {
	//matrix := [][]byte{
	//	{'1', '0'},
	//	{'0', '0'},
	//}
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	for _, line := range matrix {
		for _, ch := range line {
			fmt.Printf("| %s ", string(ch))
		}
		fmt.Print("|\n")
	}
	fmt.Println("只包含 '1' 的最大正方形的面积", maximalSquare2(matrix))
}
