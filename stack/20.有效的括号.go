/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package stack

// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		// 奇数个字符的字符串肯定不是有效的括号
		// 空字符串也不是有效的括号,所以最少2个字符
		return false
	}
	m := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	// 判断是否是右括号
	isRight := func(b byte) bool {
		_, ok := m[b]
		return !ok
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if isRight(s[i]) {
			// 右括号
			if len(stack) == 0 {
				// 栈为空,说明没有左括号匹配
				return false
			}
			// 栈顶元素出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if m[top] != s[i] {
				// 栈顶元素和当前元素不匹配
				return false
			}
		} else {
			// 左括号
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// @lc code=end
