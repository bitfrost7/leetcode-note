package dynamic_programming

import (
	"fmt"
	"leetcode-note/greedy"
)

// 给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）

// 最大k次交易
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}
	k = min(k, len(prices)/2) // k 超过 N/2 是没有意义的 因为 n天最多只能进行n/2笔交易
	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}
	return dp[k][0]
}

// 最少k次交易次数
func minKProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}
	if k > len(prices)/2 {
		return greedy.MaxProfitUnlimited(prices)
	}
	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := k; j > 0; j-- {
			dp[j][0] = min(dp[j][0], dp[j][1]-prices[i])
			dp[j][1] = min(dp[j][1], dp[j-1][0]+prices[i])
		}
	}
	return dp[k][0]
}

// 递归解法
func dfs(i, k, flag int, prices []int) int {
	if i == 0 && flag == 0 {
		return 0
	}
	if i == 0 && flag == 1 {
		return -prices[0]
	}
	if k == 0 {
		return 0
	}
	if flag == 0 {
		return max(dfs(i-1, k, 0, prices), dfs(i-1, k, 1, prices)+prices[i])
	} else {
		return max(dfs(i-1, k, 1, prices), dfs(i-1, k-1, 0, prices)-prices[i])
	}
}

// 不限次数交易，但有冷冻期
func maxProfitWithFreeze(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][3]int, len(prices))
	dp[0][0] = 0          // 不持有股票，且不处于冷冻期
	dp[0][1] = -prices[0] // 持有股票
	dp[0][2] = 0          // 不持有股票，且处于冷冻期
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])
}

// 不限次数交易，但有手续费
func maxProfitWithFee(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0                // 不持有股票
	dp[0][1] = -prices[0] - fee // 持有股票
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[len(prices)-1][0]
}

func TestMaxProfit() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(2, prices))
}
