package dynamic_programming

import (
	"fmt"
)

//https://leetcode.cn/problems/uncrossed-lines/description
//在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//
//现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
//
// nums1[i] == nums2[j]
//且绘制的直线不与任何其他连线（非水平线）相交。
//请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
//以这种方法绘制线条，并返回可以绘制的最大连线数。
//示例 1：
//输入：nums1 = [1,4,2], nums2 = [1,2,4]
//输出：2
//解释：可以画出两条不交叉的线，如上图所示。
//但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交

// 动态规划
// DP[i][j] 代表nums1前i个数和nums2前j个数的最大连接数
// 边界条件：DP[0][j] = 0 DP[i][0] = 0
// 递推公式:
//
//	DP[i][j] = {
//	   DP[i-1][j-1] + 1; 当nums1[i] == nums2[j]
//	   max(DP[i-1][j], DP[i][j-1]); 当nums1[i] != nums2[j]
//	}
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

// 空间优化至一维dp
func maxUncrossedLines1(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([]int, n+1)
	pre := 0
	for i := 1; i < m+1; i++ {
		pre = dp[0]
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[j], pre = pre+1, dp[j]
			} else {
				dp[j], pre = max(dp[j-1], dp[j]), dp[j]
			}
		}
	}
	return dp[n]
}
func TestMaxUncrossedLines() {
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}
	fmt.Printf("Input: nums1=%v,nums2=%v\n", nums1, nums2)
	fmt.Println("nums1 和 nums2 中可以绘制的最大连线数", maxUncrossedLines1(nums1, nums2))
}
