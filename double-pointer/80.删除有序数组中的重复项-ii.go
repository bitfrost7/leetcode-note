/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */
package double_pointer

// @lc code=start
func removeDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	i, j := 2, 2
	for j < len(nums) {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}

// @lc code=end
