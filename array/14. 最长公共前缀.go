package array

import "fmt"

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

// ----- Begin Test -----
func TestLongestCommonPrefix() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("strs:", strs)
	fmt.Println("longestCommonPrefix:", longestCommonPrefix(strs)) // fl

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println("strs:", strs2)
	fmt.Println("longestCommonPrefix:", longestCommonPrefix(strs2)) // ""
}

// ----- Solution Here -----

// 解法：纵向扫描
//
// 以第一个字符串为基准，从左到右逐列比较：
// - 第 i 列字符必须在所有字符串中都存在且相同，否则当前累计前缀就是答案。
//
// 复杂度：
// - 时间 O(S)，S 为所有字符串字符数总和（最坏情况扫描到最短串长度，并对每列比较所有字符串）。
// - 额外空间 O(1)。
