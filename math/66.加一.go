/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
package math

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果当前位不是9，直接加1返回
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		// 如果当前位是9，将其置为0
		digits[i] = 0
	}
	// 如果所有位都是9，数组长度加1，首位置1
	return append([]int{1}, digits...)
}

// @lc code=end
