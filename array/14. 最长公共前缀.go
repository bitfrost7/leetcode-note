package array

// description: (https://leetcode.cn/problems/longest-common-prefix/description/)
//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//
//
//示例 1：
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//示例 2：
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//

// ----- Code Here -----
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	lcp := ""
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || first[i] != strs[j][i] {
				return lcp
			}
		}
		lcp = first[:i+1]
	}
	return lcp
}

// 题解:
// 纵向或者横向扫描即可

// ----- Begin Test -----
func TestLongestCommonPrefix() string {
	a := "abcd"
	b := "ab"
	return longestCommonPrefix([]string{a, b})
}
