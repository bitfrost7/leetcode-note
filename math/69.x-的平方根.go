package math

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

// @lc code=end
