package dynamic_programming

import "fmt"

/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
*/
func longestPalindromeSubseq(s string) string {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	// 如果 s[i] == s[j]，则 dp[i][j] = dp[i+1][j-1] + 2
	// 如果 s[i] != s[j]，则 dp[i][j] = max(dp[i+1][j], dp[i][j-1]) 要知道 dp[i][j]，需要知道 dp[i+1][j] 和 dp[i][j-1]的值，所以遍历顺序是i倒序，j正序
	left, maxLength := 0, 0
	for i := len(s) - 1; i >= 0; i-- { // 从下往上遍历
		for j := i + 1; j < len(s); j++ { // 从左往右遍历
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2 // 状态转移方程
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1]) // 状态转移方程
			}
			if dp[i][j] > maxLength {
				maxLength = dp[i][j]
				left = i
			}
		}
	}
	return s[left : left+maxLength]
}

func TestLongestPalindromeSubseq() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
