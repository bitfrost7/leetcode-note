package greedy

import "fmt"

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	// 每个孩子至少分到一块钱，剩余的钱尽可能分给孩子7块钱，凑够8块钱
	money = money - children
	res := money / 7
	if res > children {
		res = children
	}

	// 剩余的钱和孩子数
	money = money - res*7
	children = children - res
	// 如果没有钱了，返回分配的次数
	if money == 0 {
		return res
	}

	// 当没有孩子时，将剩余的钱分给最后一个孩子，8+余数，res-1
	if children == 0 && money > 0 {
		return res - 1
	}
	// 当只有一个孩子时，且剩余的钱为3，因为不能凑够4块钱，将3块钱分部分给最后一个孩子，res-1
	if children == 1 && money == 3 {
		return res - 1
	}
	return res
}
func TestDistMoney() {
	money := 23
	children := 2
	fmt.Println(distMoney(money, children))
}
