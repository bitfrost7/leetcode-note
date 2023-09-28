package dynamic_programming

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	maxLen := 1
	start := 0
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
		if i < len(s)-1 && s[i] == s[i+1] {
			dp[i][i+1] = 1
			maxLen = 2
			start = i
		}
	}
	for l := 3; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			j := l + i - 1
			if s[i] == s[j] && dp[i+1][j-1] == 1 {
				dp[i][j] = 1
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
