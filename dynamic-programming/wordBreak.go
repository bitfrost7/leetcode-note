package dynamic_programming

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}
	dp := make([]bool, len(s)+1) // dp[i] 表示字符串 s 前 i 个字符组成的字符串是否能被空格拆分成若干个字典中出现的单词
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] { // dp[j] 表示 s 的前 j 个字符是否可以拆分成字典中的单词
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
