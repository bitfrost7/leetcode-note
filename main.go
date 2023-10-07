package main

import "leetcode/design"

func main() {
	stockSpanner := design.Constructor()
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	for _, price := range prices {
		println(stockSpanner.Next(price))
	}
}
