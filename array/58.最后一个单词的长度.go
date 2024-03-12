package array

import "fmt"

// description: (https://leetcode-cn.com/problems/length-of-last-word/)
// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
// 示例 1：
// 输入：s = "Hello World"
// 输出：5
func lengthOfLastWord(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			for j := i - 1; j >= 0; j-- {
				if s[j] == ' ' {
					return i - j
				}
			}
			return i + 1
		}
	}
	return 0
}

func TestLengthOfLastWord() {
	fmt.Println("Input: s = \"Hello World\"")
	fmt.Println(lengthOfLastWord("Hello World"))
}
