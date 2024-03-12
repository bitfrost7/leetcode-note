package dynamic_programming

import (
	"fmt"
)

// https://leetcode.cn/problems/longest-arithmetic-subsequence/description
// 给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。
//
// 回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ... < ik <= nums.length - 1。并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1) 的值都相同，那么序列 seq 是等差的。
// 输入：nums = [3,6,9,12]
// 输出：4
// 解释：
// 整个数组是公差为 3 的等差数列
func longestArithSeqLength(nums []int) int {
	// DP[i][d]定义为以第i个元素结尾,公差为d的最长等差子序列的长度
	dp := make([]map[int]int, len(nums))
	// 递推公式
	// d = nums[i]-nums[j]
	// DP[i][d] = DP[j][d] + 1
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j]
			if dp[i][d] == 0 {
				dp[i][d] = dp[j][d] + 1
			}
			if maxLength < dp[i][d] {
				maxLength = dp[i][d]
			}
		}
	}
	return maxLength + 1
}

func TestLongestArithSeqLength() {
	nums := []int{9, 4, 7, 2, 10}
	fmt.Printf("input: nums=%v\n", nums)
	fmt.Println("最长等差子序列的长度=", longestArithSeqLength(nums))
}
