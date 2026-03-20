package array

import "fmt"

// description: (https://leetcode-cn.com/problems/length-of-last-word/)
// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
// 示例 1：
// 输入：s = "Hello World"
// 输出：5

// ----- Code Here -----
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

// ----- Begin Test -----
func TestLengthOfLastWord() {
	fmt.Println("Input: s = \"Hello World\"")
	fmt.Println(lengthOfLastWord("Hello World"))

	fmt.Println("Input: s = \"   fly me   to   the moon  \"")
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

// ----- Solution Here -----

// 解法：从右往左扫描
//
// 从字符串末尾向左找第一个非空格字符（最后一个单词的末尾 i），
// 然后继续向左找第一个空格字符（单词前的分隔处 j），返回 i-j 即为长度。
//
// 复杂度：
// - 时间 O(n)
// - 额外空间 O(1)
