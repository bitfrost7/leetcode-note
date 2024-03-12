package dynamic_programming

import (
	"fmt"
)

// 给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。
//
// 子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。
//
// 输入：arr = [1,2,3,4], difference = 1
// 输出：4
// 解释：最长的等差子序列是 [1,2,3,4]。

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	maxLength := 1
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > maxLength {
			maxLength = dp[v]
		}
	}
	return maxLength
}
func longestSubsequence1(arr []int, difference int) int {
	if len(arr) == 1 {
		return 1
	}
	dp := make([]int, len(arr))
	dp[0] = 1
	maxLength := 1
	for i := 1; i < len(arr); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if arr[i]-arr[j] == difference && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				break
			}
		}
		if maxLength < dp[i] {
			maxLength = dp[i]
		}
	}
	return maxLength
}
func TestLongestSubsequence() {
	arr := []int{3, 0, -3, 4, -4, 7, 6}
	difference := 3
	fmt.Println("array:", arr)
	fmt.Println("difference:", difference)
	fmt.Println("最长的等差子序列长度是:", longestSubsequence1(arr, difference))
}
