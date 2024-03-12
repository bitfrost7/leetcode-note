package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/longest-common-subsequence/description
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//
// 示例 1：
//
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
// 示例 2：
//
// 输入：text1 = "abc", text2 = "abc"
// 输出：3
// 解释：最长公共子序列是 "abc" ，它的长度为 3 。
// 示例 3：
//
// 输入：text1 = "abc", text2 = "def"
// 输出：0
// 解释：两个字符串没有公共子序列，返回 0 。

// 动态规划
// DP[i][j] = LCS(text1[:i],text2[:j])
// 边界条件: DP[i][0], DP[0][j] 都为0
// 递推公式:
//
//	DP[i][j] = {
//	   DP[i-1][j-1] + 1; 当text1[i] == text2[j]
//	   max(DP[i-1][j], DP[i][j-1]); 当text1[i] != text2[j]
//	}
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// 动态规划优化到一维DP
func longestCommonSubsequence1(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([]int, n+1)
	pre := 0 // 保存dp[j]被改之前的值 相当于原来的二维dp中的对角dp值
	for i := 1; i < m+1; i++ {
		pre = dp[0]
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[j], pre = pre+1, dp[j]
			} else {
				dp[j], pre = max(dp[j-1], dp[j]), dp[j]
			}
		}
	}
	return dp[n]
}
func TestLongestCommonSubsequence() {
	text1, text2 := "abcdegggt", "acet"
	fmt.Printf("Input text1:%v,text2:%v\n", text1, text2)
	fmt.Println("最长公共子序列的长度为:", longestCommonSubsequence(text1, text2))
}
