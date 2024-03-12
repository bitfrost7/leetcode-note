/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

package string

// @lc code=start
func isIsomorphic(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}
	m1 := map[byte]byte{}
	m2 := map[byte]byte{}
	for i := 0; i < l1; i++ {
		//对于s和t中每一个字符，都有s 能够映射到t t映射到s
		x, y := s[i], t[i]
		//这一句的意思是：如果当前字符在映射集中，但是和已有的映射不相等，则返回false
		if (m1[x] != 0 && m1[x] != y) || (m2[y] != 0 && m2[y] != x) {
			return false
		}
		//如果不在映射集中，则添加映射
		m1[x] = y
		m2[y] = x
	}
	return true
}

// @lc code=end
