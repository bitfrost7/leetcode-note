package design

import "math"

type StockSpanner struct {
	stack [][2]int // [idx, price]
	idx   int      // 当前索引
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: [][2]int{{-1, math.MaxInt}}, // 初始化栈，栈底元素为最小值，防止栈空
		idx:   -1,
	}
}

func (this *StockSpanner) Next(price int) int {
	this.idx++
	for this.stack[len(this.stack)-1][1] <= price { // 栈顶元素小于当前元素，出栈
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{this.idx, price})
	return this.idx - this.stack[len(this.stack)-2][0]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
