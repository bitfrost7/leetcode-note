package dynamic_programming

// https://leetcode.cn/problems/perfect-squares/description
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
// 思路：完全平方数 很明显在于 1..sqrt(n)^2 之间, 那么问题就变成 在1..sqrt(n)中选择最少的完全平方数加起来恰好等于n
// 此时相当于 LC322. 零钱兑换 https://leetcode.cn/problems/coin-change/description/
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		minn := n + 1
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}
