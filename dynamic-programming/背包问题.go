package dynamic_programming

import "fmt"

// 问题描述:
// 有N件物品和一个容量为W的背包，第i件物品的重量为w[i],价值为v[i],求在不超过背包容量的情况下，能装入背包的最大总价值
//
// DP[i][j] 表示从下标为[0-i]的物品里任意取，放进容量为j的背包，所能取到的最大总价值
// 每件物品只有两种选择(状态), 拿或者不拿。
// 若背包容量 j 不足以放下当前物品，则保持上一状态的最大价值：DP[i][j] = DP[i-1][j]
// 否则，在放下和不放下当前物品之间选择能获得的最大价值：DP[i][j] = max(DP[i-1][j], DP[i-1][j-w[i]] + v[i])
// DP[0][j] 代表不放任何物品; DP[i][0] 代表背包容量为0

// 动态规划
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

// 一维dp
func bagProblem2(w, v []int, W int) int {
	n := len(w)
	dp := make([]int, W+1)
	for i := 1; i <= n; i++ {
		dp[0] = 0
		for j := W; j > 0; j-- {
			if j >= w[i-1] {
				dp[j] = max(dp[j-w[i-1]]+v[i-1], dp[j])
			}
		}
	}
	return dp[W]
}

func TestBagProblem() {
	weights := []int{1, 2, 3, 4, 5}
	values := []int{1, 2, 3, 4, 5}
	W := 20
	for i := range weights {
		fmt.Printf("物品%d:重量%d价值%d\n", i, weights[i], values[i])
	}
	fmt.Println("能装入背包的最大总价值", bagProblem2(weights, values, W))
}
