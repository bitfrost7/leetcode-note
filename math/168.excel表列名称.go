package math

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	res := []byte{}
	for columnNumber > 0 {
		columnNumber--
		res = append(res, byte(columnNumber%26)+'A')
		columnNumber /= 26
	}
	return func(s []byte) string {
		l, r := 0, len(s)-1
		for l < r {
			s[l], s[r] = s[r], s[l]
			l++
			r--
		}
		return string(s)
	}(res)
}

// @lc code=end
