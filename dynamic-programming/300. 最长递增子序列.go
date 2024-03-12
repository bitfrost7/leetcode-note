package dynamic_programming

import "fmt"

/*
https://leetcode.cn/problems/longest-increasing-subsequence/description
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1
*/

// 动态规划
// 设dp[i] = 前i个数字中以第i个数字结尾最长严格递增子序列的长度
// DP[i] = max(枚举之前每个数字，如果比他大，则dp[i] = DP[i-n]+1)
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	maxLength := 0
	for _, l := range dp {
		if maxLength < l {
			maxLength = l
		}
	}
	return maxLength
}

func TestLengthOfLIS() {
	nums := []int{4, 10, 4, 3, 8, 9, 4, 4, 4}
	fmt.Println("input:", nums)
	fmt.Println("最长递增子序列长度:", lengthOfLIS(nums))
}
