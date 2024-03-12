package dynamic_programming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv
// 给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1：
//
// 输入：k = 2, prices = [2,4,1]
// 输出：2
// 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
// 示例 2：
//
// 输入：k = 2, prices = [3,2,6,5,0,3]
// 输出：7
// 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
// 随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
func maxProfitIV(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}
	k = min(len(prices)/2, k)
	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][0] = 0          //代表第i次交易 不持有股票
		dp[i][1] = -prices[0] //代表第i次交易 持有股票
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])   // 要么是之前就不持有股票 要么之前持有股票但是卖出了
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i]) // 要么之前持有股票 要么上次交易完 又买入了
		}
	}
	return dp[k][0]
}
