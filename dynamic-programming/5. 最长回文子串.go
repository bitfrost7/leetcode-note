package dynamic_programming

import "fmt"

/*
https://leetcode.cn/problems/longest-palindromic-substring/description
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
examples:
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案
*/

//1. 动态规划解法
//DP[i][j] 代表第i个字符到第j个字符是否是回文字符串
//递推公式： DP[i][j] = DP[i+1][j-1] ^ s[i] == s[j]
//如果一个字符串是回文字符串，那么该字符串的子串也应该是回文字符串
//边界情况1:一个字符一定是回文字符串 2:两个字符 如果相等那就是回文字符串
//递推 时间复杂度O(n^2) 空间复杂度 O(n^2)

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	maxLen, start := 1, 0
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if i < len(s)-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}
	for l := 3; l <= len(s); l++ {
		for i := 0; i < len(s)-l+1; i++ {
			j := i + l - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if l > maxLen {
					start = i
					maxLen = l
				}
			}
		}
	}
	return s[start : start+maxLen]
}

// 2.中心扩展算法
func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, end := 0, 0
	for i := range s {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, start, end int) (int, int) {
	for start >= 0 && end <= len(s)-1 {
		if s[start] == s[end] {
			start -= 1
			end += 1
		} else {
			break
		}
	}
	return start + 1, end - 1
}

func TestLongestPalindrome() {
	s := "cbbd"
	fmt.Println("origin:", s)
	fmt.Println("palindrome:", longestPalindrome(s))
	fmt.Println("palindrome2:", longestPalindrome2(s))
}
