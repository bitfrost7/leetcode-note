package dynamic_programming

import "fmt"

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) // dp[i][j] 表示 word1 的前 i 个字母和 word2 的前 j 个字母之间的编辑距离
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // word2 为空，将 word1 的全部删除
	}
	for j := range dp[0] {
		dp[0][j] = j // word1 为空，插入 word2 的全部字符
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown += 1
			}
			dp[i][j] = min(left, min(down, leftDown))
		}
	}
	return dp[m][n]
}

func TestminDistance() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(word1, word2)
	fmt.Println(minDistance(word1, word2))
}
