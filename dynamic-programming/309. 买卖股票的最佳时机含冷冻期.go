package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown
// 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1:
//
// 输入: prices = [1,2,3,0,2]
// 输出: 3
// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
// 示例 2:
//
// 输入: prices = [1]
// 输出: 0

func maxProfitWithCoolDown(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0] // 持有股票
	dp[0][1] = 0          // 不持有股票 不在冷冻期
	dp[0][2] = 0          // 不持有股票 在冷冻期
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i]
	}
	return max(dp[len(dp)-1][1], dp[len(dp)-1][2])
}

// 空间优化
func maxProfitWithCoolDown1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var (
		d0 = -prices[0]
		d1 = 0
		d2 = 0
	)
	for i := 1; i < len(prices); i++ {
		d0, d1, d2 = max(d0, d1-prices[i]), max(d1, d2), d0+prices[i]
	}
	return max(d1, d2)
}
func TestMaxProfitWithCoolDown() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Printf("Input: %v\n", prices)
	fmt.Println("含冷冻期的最大利润", maxProfitWithCoolDown(prices))
}
