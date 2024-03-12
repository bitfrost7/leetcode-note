package enumerate

import "fmt"

// https://leetcode.cn/problems/two-sum/description/
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	length := len(nums)
	m := make(map[int]int, length)
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{m[target-v], i}
		} else {
			m[v] = i
		}
	}
	return nil
}
func twoSum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
func TestTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target)
	fmt.Println(twoSum1(nums, target))
}
