package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/wildcard-matching/description/
//给你一个输入字符串 (s) 和一个字符模式 (p) ，请你实现一个支持 '?' 和 '*' 匹配规则的通配符匹配：
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符序列（包括空字符序列）。
//判定匹配成功的充要条件是：字符模式必须能够 完全匹配 输入字符串（而不是部分匹配）。

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func TestIsMatch() {
	s := "aa"
	p := "*"
	fmt.Printf("Input: s=%v,p=%v\n", s, p)
	fmt.Printf("是否能匹配:%v\n", isMatch(s, p))
}
