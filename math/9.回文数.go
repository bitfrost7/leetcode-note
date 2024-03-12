/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
package math

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		// 负数或者10的倍数
		return false
	}
	left, right := x, 0
	for left > right {
		right = right*10 + left%10
		left /= 10
	}
	if left == right || left == right/10 {
		return true
	}
	return false
}

// @lc code=end
