/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */
package double_pointer

// @lc code=start
func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}
	return i + 1
}

// @lc code=end
