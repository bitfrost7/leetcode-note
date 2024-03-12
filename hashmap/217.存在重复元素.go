/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

package hashmap

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

// @lc code=end
