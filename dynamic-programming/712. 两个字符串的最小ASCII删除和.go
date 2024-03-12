package dynamic_programming

import "fmt"

/*
https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/description
给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。
示例 1:

输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
示例 2:

输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
*/

// 动态规划
// 设dp[i][j]为s1前i个字符和s2前j个字符相等所需的最小ascii和
// if c1 != c2 DP[i][j] = min(DP[i-1][j]+int(c1), DP[i][j-1]+int(c2))
// if c1 == c2 DP[i][j] = DP[i-1][j-1]
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for j := range dp[0] {
		if j > 0 {
			dp[0][j] = dp[0][j-1] + int(s2[j-1])
		}
	}
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(c1), dp[i+1][j]+int(c2))
			}
		}
	}
	return dp[m][n]
}

// 空间优化
func minimumDeleteSum1(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([]int, n+1)
	pre := 0
	for i := range s2 {
		dp[i+1] = dp[i] + int(s2[i])
	}
	for i := 1; i < m+1; i++ {
		pre = dp[0]
		dp[0] += int(s1[i-1])
		for j := 1; j < n+1; j++ {
			tmp := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = pre
			} else {
				dp[j] = min(dp[j]+int(s1[i-1]), dp[j-1]+int(s2[j-1]))
			}
			pre = tmp
		}
	}
	return dp[n]
}

func minByte(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func TestMinimumDeleteSum() {
	s1 := "sea"
	s2 := "eat"
	fmt.Printf("s1:%s s2:%s\n", s1, s2)
	fmt.Println("使两个字符串相等所需删除字符的 ASCII 值的最小和:", minimumDeleteSum(s1, s2))
}
