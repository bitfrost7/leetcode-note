/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */
package hashmap

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	for i, v := range nums {
		if pos, ok := m[v]; ok {
			if abs(pos, i) <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}

// @lc code=end
