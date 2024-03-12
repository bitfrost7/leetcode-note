package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/longest-palindromic-subsequence/description/
// 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
//
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
// 示例 1：
// 输入：s = "bbbab"
// 输出：4
// 解释：一个可能的最长回文子序列为 "bbbb" 。
// 示例 2：
// 输入：s = "cbbd"
// 输出：2
// 解释：一个可能的最长回文子序列为 "bb" 。

// 定义dp[i][j]为字符串中i到j之间的最长回文子序列的长度
// 递推公式为：DP[i][j] = DP[i+1][j-1] + 2 或者 max(DP[i+1][j], DP[i][j-1]) 取决于 s[i] 是否等于 s[j]
// 因为dp[i][j]是从dp[i+1][j]和dp[i][j-1])转移而来所以i从上往下遍历, j从左往右遍历
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- { // 从下往上遍历
		for j := i + 1; j < len(s); j++ { // 从左往右遍历
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2 // 状态转移方程
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1]) // 状态转移方程
			}
		}
	}
	return dp[0][len(s)-1]
}

func TestLongestPalindromeSubseq() {
	s := "cbbd"
	fmt.Println("Input:", s)
	fmt.Println("最长回文子序列", longestPalindromeSubseq(s))
}
