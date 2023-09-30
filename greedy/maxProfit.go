package greedy

import "fmt"

func maxProfit(prices []int) int {
	minPrice, profit := prices[0], 0
	for _, v := range prices {
		minPrice = min(minPrice, v)
		profit = max(profit, v-minPrice)
	}
	return profit
}

func TestMaxProfit() {
	prices := []int{7, 4, 5, 1, 10}
	fmt.Println(maxProfit(prices))
}
