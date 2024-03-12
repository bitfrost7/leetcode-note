package math

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	major, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count++
		} else {
			count--
		}
	}
	return major
}

// @lc code=end
