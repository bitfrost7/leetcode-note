package dynamic_programming

import (
	"fmt"
)

/*
https://leetcode.cn/problems/edit-distance/description
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/

// 动态规划
// 设dp[i][j]为将word1前i个字符转换为word2前j个字符所使用的最少操作数
// 递推公式：
// DP[i][j] = min(DP[i-1][j]+1,DP[i][j-1]+1,DP[i-1][j-1]+int(word1[i]!=word2[j])
// 解释：
// 举个例子，word1 = "abcde", word2 = "fgh",我们现在算这俩字符串的编辑距离，就是找从word1，最少多少步，能变成word2？那就有三种方式：
// 知道"abcd"变成"fgh"多少步（假设X步），那么从"abcde"到"fgh"就是"abcde"->"abcd"->"fgh"。（一次删除，加X步，总共X+1步）
// 知道"abcde"变成“fg”多少步（假设Y步），那么从"abcde"到"fgh"就是"abcde"->"fg"->"fgh"。（先Y步，再一次添加，加X步，总共Y+1步）
// 知道"abcd"变成“fg”多少步（假设Z步），那么从"abcde"到"fgh"就是"abcde"->"fge"->"fgh"。（先不管最后一个字符，把前面的先变好，用了Z步，然后把最后一个字符给替换了。这里如果最后一个字符碰巧就一样，那就不用替换，省了一步）

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) // DP[i][j] 表示 word1 的前 i 个字母和 word2 的前 j 个字母之间的编辑距离
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // word2 为空，将 word1 的全部删除
	}
	for j := range dp[0] {
		dp[0][j] = j // word1 为空，插入 word2 的全部字符
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			x := dp[i-1][j] + 1
			y := dp[i][j-1] + 1
			z := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				z += 1
			}
			dp[i][j] = min(x, min(y, z))
		}
	}
	return dp[m][n]
}

// 空间优化
func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([]int, m+1)
	z := 0
	for i := range dp {
		dp[i] = i
	}
	for i := 1; i < n+1; i++ {
		z = dp[0]
		dp[0] = i
		for j := 1; j < m+1; j++ {
			if word2[i-1] != word1[j-1] {
				z += 1
			}
			dp[j], z = min(min(dp[j]+1, dp[j-1]+1), z), dp[j]
		}
	}
	return dp[m]
}

func TestMinDistance() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println("word1:", word1, "word2:", word2)
	fmt.Println("最小编辑距离：", minDistance1(word1, word2))
}
