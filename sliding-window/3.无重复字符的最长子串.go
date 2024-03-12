package sliding_window

import "fmt"

// description: 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	maxSubstr, rk := 0, -1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m := map[byte]int{}
	for i := range s {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < len(s) && m[s[rk+1]] == 0 {
			m[s[rk+1]] = 1
			rk++
		}
		maxSubstr = max(maxSubstr, rk+1-i)
	}
	return maxSubstr
}

func TestLengthOfLongestSubstring() {
	s := "abcabcbb"
	fmt.Println("Input:", s)
	fmt.Println(lengthOfLongestSubstring(s))
}
