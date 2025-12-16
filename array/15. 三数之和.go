package array

import (
	"fmt"
	"sort"
)

// description: (https://leetcode.cn/problems/3sum/description)
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

// ----- Code Here -----

func threeSum(nums []int) [][]int {
	// 1. 排序 O(N log N)
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	// 2. 遍历第一个数 a (固定) O(N)
	for i := 0; i < n-2; i++ {
		// 如果固定的 a 已经大于 0，则不可能找到和为 0 的三元组
		if nums[i] > 0 {
			break
		}
		
		// 去重 a：跳过重复的 a。i > 0 确保第一个 a 被处理。
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 3. 双指针 L (b) 和 R (c)
		L := i + 1
		R := n - 1

		for L < R {
			sum := nums[i] + nums[L] + nums[R]

			if sum == 0 {
				// 找到一个有效解
				res = append(res, []int{nums[i], nums[L], nums[R]})

				L++
				R--
				
				// 跳过重复的 L (b)：新 L 指向的值如果与前一个值相同，则继续移动
				for L < R && nums[L] == nums[L-1] {
					L++
				}
				// 跳过重复的 R (c)：新 R 指向的值如果与后一个值相同，则继续移动
				for L < R && nums[R] == nums[R+1] {
					R--
				}

			} else if sum < 0 {
				// 和太小，需要更大的数，L 向右移动，并跳过重复的 L
				L++
				for L < R && nums[L] == nums[L-1] {
					L++
				}
			} else { // sum > 0
				// 和太大，需要更小的数，R 向左移动，并跳过重复的 R
				R--
				for L < R && nums[R] == nums[R+1] {
					R--
				}
			}
		}
	}
	return res
}

// ----- Begin Test -----

func TestThreeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("nums:", nums)
	fmt.Println("threeSum is :", threeSum(nums))
}

// ----- Solution Here -----
