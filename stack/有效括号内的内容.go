package stack

import "fmt"

/*
题目，输出每对有效括号内的内容；

输入："(1+1)*2" 输出：1+1

输入： "((1+2)*(3+4))+2"  输出：1+2, 3+4, (1+2)*(3+4)

*/

func validContent(content string) []string {
	res := make([]string, 0) // 存储结果
	stk := make([]int, 0)    // 记录左括号的位置
	for i, ch := range content {
		if ch == '(' {
			stk = append(stk, i)
		} else if ch == ')' {
			if len(stk) > 0 {
				left := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				right := i
				res = append(res, content[left+1:right])
			}
		}
	}
	return res
}

func TestValidContent() {
	contents := []string{
		"(1+1)*2",
		"((1+2)*(3+(4+5)))+(6+7)",
		"))eee(fff)(f",
		"",
		"(qq(rrr(ff)))",
		"((x+y)*z)+((a+(b*c))*(d+e))",
	}
	for _, content := range contents {
		res := validContent(content)
		fmt.Println(res)
	}
}
