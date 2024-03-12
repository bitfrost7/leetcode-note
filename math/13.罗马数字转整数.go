/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
package math

// @lc code=start
func romanToInt(s string) int {
	var res int
	var m = map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && m[s[i]] < m[s[i+1]] {
			res -= m[s[i]]
		} else {
			res += m[s[i]]
		}
	}
	return res
}

// @lc code=end
