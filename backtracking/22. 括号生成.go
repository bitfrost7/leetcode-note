package backtracking

import "fmt"

// description: (https://leetcode-cn.com/problems/generate-parentheses/)
// 给定 n 代表生成括号的对数，请你写出一个函数，
// 使其能够生成所有可能的并且有效的括号组合。
// 例如
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]

func generateParenthesis(n int) []string {
	var res []string
	dfs(n, 0, 0, "", &res)
	return res
}

// n 代表生成括号的对数 left right 分别代表左右括号的数量 path 代表当前路径 res 代表结果集
func dfs(n, left, right int, path string, res *[]string) {
	if left == n && right == n {
		*res = append(*res, path)
		return
	}
	// 如果左括号数量小于等于右括号数量，此时只能添加左括号
	if left <= right {
		dfs(n, left+1, right, path+"(", res)
	} else {
		// 如果左括号数量大于右括号数量，此时可以添加左括号也可以添加右括号
		if left < n {
			dfs(n, left+1, right, path+"(", res)
		}
		dfs(n, left, right+1, path+")", res)
	}
}

func TestGenerateParenthesis() {
	fmt.Println("Input: n = 3")
	fmt.Println(generateParenthesis(3))
}
