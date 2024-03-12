package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/coin-change-ii/description/
// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//
// 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
// 假设每一种面额的硬币有无限个。
//
// 题目数据保证结果符合 32 位带符号整数。
func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, coin := range coins {
		for i := 1; i < amount+1; i++ {
			if i-coin >= 0 {
				f[i] += f[i-coin]
			}
		}
	}
	return f[amount]
}

// 为什么先遍历硬币？
// 因为遍历金钱 上一个状态 包含所有硬币种类 转移到下一个状态 也是所有硬币种类 相当于笛卡尔积 结果和硬币的顺序有关 是排列 而非 组合

func TestChange() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Printf("Input: amount=%d coins=%v", amount, coins)
	fmt.Println("可以凑成总金额的硬币组合数", change(amount, coins))
}
