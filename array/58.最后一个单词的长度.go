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

// 题解:
// 这道题关键在于，理解数组索引在题中的含义，i是第一个不为空格的位置索引，j 为第一个为空格的位置索引，相减即为第一个单词的长度
// 所以 引申出一个定义， 索引相减则为索引左开右闭/左闭右开之间的距离。
// 同时需要这一特殊情况，倒数第一个字符是空格，倒数最后一个字符是空格，两种特殊情况

func TestLengthOfLastWord() {
	fmt.Println("Input: s = \"Hello World\"")
	fmt.Println(lengthOfLastWord("Hello World"))
}
