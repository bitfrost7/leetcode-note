package dynamic_programming

/*
https://leetcode.cn/problems/number-of-longest-increasing-subsequence/description/
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。

注意 这个数列必须是 严格 递增的。

示例 1:

输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:

输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
*/

// 动态规划：
// 定义cnt[i] 为以第i个数字结尾的最长递增子序列的个数
// 考虑当num[i]>nums[j]时，num[i]可以接在nums[j]结尾的最长递增子序列的后面
// 若dp[i]小于dp[j]+1，说明找到了一个更长的递增子序列，更新dp[i]为dp[j]+1，并将相应的子序列数量cnt[i]设为cnt[j]。
// 若dp[i]等于dp[j]+1，则说明找到了另一个同样长度的最长递增子序列，因此增加对应长度子序列的计数，即cnt[i] += cnt[j]。
func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	dp[0] = 1
	cnt[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}
	}
	maxLength := 0
	for _, l := range dp {
		if maxLength < l {
			maxLength = l
		}
	}
	ans := 0
	for i, c := range cnt {
		if dp[i] == maxLength {
			ans += c
		}
	}
	return ans
}
