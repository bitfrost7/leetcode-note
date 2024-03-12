package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii
// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1:
//
// 输入：prices = [3,3,5,0,0,3,1,4]
// 输出：6
// 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//
//	随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//
// 示例 2：
//
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//
//	注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//	因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
// 示例 3：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
// 示例 4：
//
// 输入：prices = [1]
// 输出：0
func maxProfitIII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2][2]int, len(prices))
	dp[0][0][0] = -prices[0] // 只进行过一次买操作,持有股票
	dp[0][0][1] = 0          // 完成了一笔交易，没有持有股票
	dp[0][1][0] = -prices[0] // 在完成了一笔交易的前提下，进行了第二次买操作，持有股票
	dp[0][1][1] = 0          // 完成了二次交易，没有持有股票
	for i := 1; i < len(prices); i++ {
		dp[i][0][0] = max(dp[i-1][0][0], -prices[i])
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]+prices[i])
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][1]-prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][1][0]+prices[i])
	}
	return max(0, max(dp[len(prices)-1][0][1], dp[len(prices)-1][1][1]))
}

// 空间优化版
func maxProfitIII1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy1 := -prices[0] // 只进行过一次买操作,持有股票
	sell1 := 0         // 完成了一笔交易，没有持有股票
	buy2 := -prices[0] // 在完成了一笔交易的前提下，进行了第二次买操作，持有股票
	sell2 := 0         // 完成了二次交易，没有持有股票
	for i := 1; i < len(prices); i++ {
		buy1, sell1, buy2, sell2 = max(buy1, -prices[i]), max(sell1, buy1+prices[i]), max(buy2, sell1-prices[i]), max(sell2, buy2+prices[i])
	}
	return max(0, max(sell1, sell2))
}
func TestMaxProfitIII() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Printf("Input: %v\n", prices)
	fmt.Printf("最多可以完成两笔交易所能获取的最大利润为:%d\n", maxProfitIII(prices))
}
