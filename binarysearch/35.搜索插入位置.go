package binarysearch

import "fmt"

// description: (https://leetcode-cn.com/problems/search-insert-position/)
// 给定一个「排序」数组和一个目标值，在数组中找到「目标值」，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// TestSearchInsert 测试搜索插入位置
func TestSearchInsert() {
	nums := []int{1, 3, 5, 6}
	target := 5
	// target := 2
	// target := 7
	// target := 0
	fmt.Println("Input: nums = [1,3,5,6], target = 5")
	fmt.Println(searchInsert(nums, target))
}
