package dynamic_programming

import "fmt"

/*
https://leetcode.cn/problems/climbing-stairs/description
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
*/

// 动态规划
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 优化空间
func climbStairs2(n int) int {
	p, q := 1, 1
	for i := 2; i < n+1; i++ {
		p, q = q, p+q
	}
	return q
}

func TestClimbStairs() {
	N := 3
	fmt.Println("input: N=", N)
	fmt.Println("climb_stairs:", climbStairs2(N))
}
