/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */
package double_pointer

// @lc code=start
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	isAlphanumeric := func(b byte) bool {
		return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
	}
	isSame := func(b1, b2 byte) bool {
		if b1 >= 'A' && b1 <= 'Z' {
			b1 += 32
		}
		if b2 >= 'A' && b2 <= 'Z' {
			b2 += 32
		}
		return b1 == b2
	}
	for l < r {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		if !isSame(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

// @lc code=end
