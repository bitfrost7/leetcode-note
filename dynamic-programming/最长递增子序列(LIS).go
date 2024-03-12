package dynamic_programming

import (
	"fmt"
	"sort"
)

/*
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

// LIS 最长递增子序列的解法一般有两种：动态规划和贪心+二分查找

// 动态规划模版
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := 0
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// 贪心+二分查找模版
//
// d[i]表示长度为i的最长递增子序列的末尾元素的最小值, 可以贪心的认为只要让d[i]尽可能的小，后面的元素才能接的更长，最终的LIS的最大长度就是d的长度
//
// 遍历nums数组，在d中二分搜索第一个大于等于num[i]的地方
// 如果找到的位置等于len(d)说明nums[i] > max(d) ，直接追加到d
// 否则，将d中该位置替换为nums[i]

func lengthOfLIS2(nums []int) int {
	d := make([]int, 0)
	for _, v := range nums {
		if i := sort.SearchInts(d, v); i < len(d) {
			d[i] = v
		} else {
			d = append(d, v)
		}
	}
	return len(d)
}

// 变种题目：最长严格递增子序列的个数
func findNumOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	maxLength := 1
	ans := 1
	for i := range nums {
		dp[i], cnt[i] = 1, 1
		for j := range nums[:i] {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxLength {
			ans += cnt[i]
		}
	}
	return ans
}

// 贪心+前缀和
// 首先定义d[i]为长度为i+1的LIS的所有结尾数的选择 比如d[2] = {5,3,1} 当LIS长度为3时,可能结尾有5,3,1,三种选择
// 在d的基础之上,我们定义cnt[i][j]为d[i][j]情况下的所有LIS个数,
// 我们依旧是贪心的认为在当前i个数字中,LIS增长的越慢,后面能追加更多的数字.
// 所以在遍历nums的时候, 将num插入到d中第一个大于等于num的位置(d是单调递增的)
// cnt[i][j]的值应该是cnt[i-1]中所有结尾数字小于d[i][j]的所有cnt之和
// 我们在cnt数组中直接存入前缀和简化运算
func findNumOfLIS1(nums []int) int {
	var (
		d   [][]int
		cnt [][]int
	)
	for _, v := range nums {
		// 默认有一个
		c := 1
		// 先查找v应该插入哪个长度的LIS之后
		i := sort.Search(len(d), func(i int) bool { return d[i][len(d[i])-1] >= v })
		if i > 0 {
			// 如果长度不为1,cnt的值应该算上之前长度的所有小于自己的所有长度之和
			k := sort.Search(len(d[i-1]), func(k int) bool { return d[i-1][k] < v })
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}
		// 如果是新的长度 直接追加
		if i == len(d) {
			d = append(d, []int{v})
			cnt = append(cnt, []int{0, c}) // 类似前缀和的数组 需要前面补0
		} else {
			d[i] = append(d[i], v)                           // 追加到d中长度为i+1的可能数组中
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1]+c) // 增加前缀和
		}
	}
	return cnt[len(cnt)-1][len(cnt[len(cnt)-1])-1] // 直接返回最后一个长度(最大长度)的最后的一个元素(这个长度下所有可能的和)
}

// 变种题目：最长不严格递增子序列的长度
// 例如：[10, 9, 2, 2, 5, 3, 7, 101, 18] 最长不严格递增子序列长度为5

func lengthOfLIS3(nums []int) int {
	d := make([]int, 0)
	for _, v := range nums {
		if i := sort.SearchInts(d, v+1); i < len(d) {
			d[i] = v
		} else {
			d = append(d, v)
		}
	}
	return len(d)
}

func TestLIS() {
	nums := []int{10, 9, 2, 2, 5, 5, 7, 101, 18}
	fmt.Println("input:", nums)
	fmt.Println("最长递增子序列长度:", lengthOfLIS2(nums))
	fmt.Println("最长递增子序列个数为:", findNumOfLIS1(nums))
	fmt.Println("最长不严格递增子序列长度:", lengthOfLIS3(nums))
}
