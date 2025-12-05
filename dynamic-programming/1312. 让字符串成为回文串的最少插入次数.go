package dynamic_programming

import (
	"fmt"
)

// https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/
//给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
//
//请你返回让 s 成为回文串的 最少操作次数 。
//
//「回文串」是正读和反读都相同的字符串

// 示例 1：
//
// 输入：s = "zzazz"
// 输出：0
// 解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。
// 示例 2：
//
// 输入：s = "mbadm"
// 输出：2
// 解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。
// 示例 3：
//
// 输入：s = "leetcode"
// 输出：5
// 解释：插入 5 个字符后字符串变为 "leetcodocteel"
func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	minInsert := n
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == s[n-j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
			if i+j == n-1 || i+j == n || i == n {
				minInsert = min(minInsert, dp[i][j])
			}
		}
	}
	return minInsert
}

// 空间优化
// [0 1 2 3 4 5]
// [1 1 2 3 4 4]
// [2 2 3 4 4 5]
// [3 3 4 4 5 6]
// [4 4 4 5 6 7]
// [5 5 5 6 7 7]
func minInsertions1(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	minInsert, pre := n, 0
	for i := range dp {
		dp[i] = i
	}
	for i := 1; i < n+1; i++ {
		pre = dp[0]
		dp[0] = i
		for j := 1; j < n+1; j++ {
			if s[i-1] == s[n-j] {
				dp[j], pre = pre, dp[j]
			} else {
				dp[j], pre = min(dp[j], dp[j-1])+1, dp[j]
			}
			if i+j == n-1 || i+j == n || i == n {
				minInsert = min(minInsert, dp[j])
			}
		}
	}
	return minInsert
}
func TestMinInsertions() {
	s := "zzazz"
	fmt.Printf("Input: %v\n", s)
	fmt.Printf("让字符串成为回文串的最少操作次数:%d", minInsertions1(s))
}
