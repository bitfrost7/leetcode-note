package dynamic_programming

import (
	"fmt"
	"slices"
)

// https://leetcode.cn/problems/delete-and-earn/description/
// 给你一个整数数组 nums ，你可以对它进行一些操作。
//
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
//
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
// 示例 1：
//
// 输入：nums = [3,4,2]
// 输出：6
// 解释：
// 删除 4 获得 4 个点数，因此 3 也被删除。
// 之后，删除 2 获得 2 个点数。总共获得 6 个点数。
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	arr := make([]int, slices.Max(nums)+1)
	for _, v := range nums {
		arr[v] += v
	}
	return rob(arr)
}
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	curr, prev := max(nums[0], nums[1]), nums[0]
	for i := 2; i < len(nums); i++ {
		curr, prev = max(curr, prev+nums[i]), curr
	}
	return curr
}

func TestDeleteAndEarn() {
	nums := []int{3, 4, 2}
	fmt.Println("Input: nums=", nums)
	fmt.Println("Output:", deleteAndEarn(nums))
}
