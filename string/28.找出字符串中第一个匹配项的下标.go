/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */
package string

import "strings"

// @lc code=start
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// @lc code=end
