package dynamic_programming

// https://leetcode.cn/problems/coin-change/description/
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		minn := amount + 1 //硬币的面额最少为1 所以f[i]最大为amount+1 用来表示上限
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			minn = min(minn, f[i-coins[j]])
		}
		f[i] = minn + 1
	}
	if f[amount] > amount {
		return -1
	}
	return f[amount]
}
