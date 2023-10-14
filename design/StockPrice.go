package design

import (
	"fmt"
	"math"
)

type StockPrice struct {
	prices map[int]int
	curr   int
	min    int
	max    int
}

func StockPriceConstructor() StockPrice {
	return StockPrice{
		prices: make(map[int]int),
		min:    math.MaxInt,
		max:    0,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.prices[timestamp] = price
	if timestamp >= this.curr {
		this.curr = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.prices[this.curr]
}

func (this *StockPrice) Maximum() int {
	this.max = 0
	for _, v := range this.prices {
		if v > this.max {
			this.max = v
		}
	}
	return this.max
}

func (this *StockPrice) Minimum() int {
	this.min = math.MaxInt
	for _, v := range this.prices {
		if v < this.min {
			this.min = v
		}
	}
	return this.min
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
func TestStockPrice() {
	commands := []string{"StockPrice", "update", "maximum", "current", "minimum", "maximum", "maximum", "maximum", "minimum", "minimum", "maximum", "update", "maximum", "minimum", "update", "maximum", "minimum", "current", "maximum", "update", "minimum", "maximum", "update", "maximum", "maximum", "current", "update", "current", "minimum", "update", "update", "minimum", "minimum", "update", "current", "update", "maximum", "update", "minimum"}
	values := [][]interface{}{{}, {38, 2308}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {47, 7876}, {}, {}, {58, 1866}, {}, {}, {}, {}, {43, 121}, {}, {}, {40, 5339}, {}, {}, {}, {32, 5339}, {}, {}, {43, 6414}, {49, 9369}, {}, {}, {36, 3192}, {}, {48, 1006}, {}, {53, 8013}, {}}

	var stockPrice StockPrice
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		value := values[i]

		switch command {
		case "StockPrice":
			stockPrice = StockPriceConstructor()
		case "update":
			timestamp := value[0].(int)
			price := value[1].(int)
			stockPrice.Update(timestamp, price)
		case "maximum":
			fmt.Println(stockPrice.Maximum())
		case "current":
			fmt.Println(stockPrice.Current())
		case "minimum":
			fmt.Println(stockPrice.Minimum())
		}
	}
}
