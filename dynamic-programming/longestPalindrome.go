package dynamic_programming

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// dp[i][j] 代表第i个字符到第j个字符是否是回文字符串
	// 递推公式： dp[i][j] = dp[i+1][j-1] ^ s[i] == s[j]
	// 如果一个字符串是回文字符串，那么该字符串的子串也应该是回文字符串
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	maxLen := 1
	start := 0
	// 边界情况：1 一个字符的子串一定是回文字符串 2 两个字符的子串 如果相等那就是回文字符串
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if i < len(s)-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
			maxLen = 2
			start = i
		}
	}
	// 递推 时间复杂度O(n^2) 空间复杂度 O(n^2)
	for l := 3; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			j := l + i - 1
			if s[i] == s[j] && dp[i+1][j-1] == true {
				dp[i][j] = true
				maxLen = l
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}
func TestLongestPalindrome(s string) string {
	return longestPalindrome(s)
}
