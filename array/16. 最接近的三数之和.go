package array

import (
	"fmt"
	"sort"
)

// description: (https://leetcode.cn/problems/3sum-closest/description/)
//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
//返回这三个数的和。
//
//假定每组输入只存在恰好一个解。

// ----- Code Here -----

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	best := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum, target) < abs(best, target) {
				best = sum
			}
			if sum == target {
				return sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return best
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// ----- Begin Test -----

func TestThreeSumClosest() {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 2, 1, -4}, target: 1, want: 2},
		{nums: []int{0, 0, 0}, target: 1, want: 0},
		{nums: []int{1, 1, 1, 0}, target: -100, want: 2},
		{nums: []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, target: -2, want: -2},
	}

	for i, tc := range tests {
		got := threeSumClosest(append([]int(nil), tc.nums...), tc.target)
		fmt.Printf("case %d: nums=%v target=%d got=%d want=%d\n", i, tc.nums, tc.target, got, tc.want)
	}
}

// ----- Solution Here -----

// 解法：排序 + 双指针
//
// 1) 先对 nums 排序。
// 2) 枚举第一个数 nums[i]，然后在区间 [i+1, n-1] 内用双指针 l、r 逼近。
//    sum = nums[i] + nums[l] + nums[r]
//    - 若 sum 更接近 target，则更新 best
//    - 若 sum < target，说明需要更大的和，l++
//    - 若 sum > target，说明需要更小的和，r--
//    - 若 sum == target，直接返回（最接近）
//
// 复杂度：
// - 排序 O(n log n)
// - 外层枚举 O(n)，内层双指针 O(n)，总计 O(n^2)
// - 额外空间 O(1)（不计排序栈开销）
