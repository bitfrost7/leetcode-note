package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/unique-binary-search-trees/description
// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

// 动态规划  设dp(n)为 n 个节点组成二叉搜索树的种数
// 递推公式:
// 假设从1..n中选择一个i节点 [1..(i-1)]作为左子树 [i+1..n]作为右子树, 可以递归的求解左子树二叉搜索树的种数和右子树二叉搜索树的种数 明显子问题可以复用
// 以i为根节点 DP(n) = ∑(i=1..n) DP(i-1)*DP(n-i)
// 解释
// 1. 根为i的所有二叉搜索树的集合是左子树集合和右子树集合的笛卡尔积
// 2. 二叉搜索树组合和节点的值无关 只和节点的个数有关
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func TestNumTrees() {
	n := 3
	fmt.Printf("Input: n=%d\n", n)
	fmt.Println("组成二叉搜索树的种数", numTrees(n))
}
