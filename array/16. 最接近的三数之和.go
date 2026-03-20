package array

import "sort"

// description: (https://leetcode.cn/problems/3sum-closest/description/)
//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
//返回这三个数的和。
//
//假定每组输入只存在恰好一个解。

// ----- Code Here -----

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minSum := 0
	n := len(nums)
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nunms[j] + nums[k]
			minSum = min(minSum, abs(sum, target))
		}
	}
}

func abs(a, b int) {
	if a > b {
		return a - b
	}
	return b - a
}

// ----- Begin Test -----

// ----- Solution Here -----
