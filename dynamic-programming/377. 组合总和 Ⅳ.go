package dynamic_programming

import "fmt"

// https://leetcode.cn/problems/combination-sum-iv/description/
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
// 题目数据保证答案符合 32 位整数范围。
func combinationSum4(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i < target+1; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				f[i] += f[i-num]
			}
		}
	}
	return f[target]
}

func TestCombinationSum4() {
	target := 4
	nums := []int{1, 2, 3}
	fmt.Printf("Input: target=%d, nums=%v", target, nums)
	fmt.Println("总和为 target 的元素组合的个数", combinationSum4(nums, target))
}
