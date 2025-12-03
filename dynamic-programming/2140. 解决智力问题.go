package dynamic_programming

import (
	"fmt"
	. "leetcode-note/helpers"
)

// https://leetcode.cn/problems/solving-questions-with-brainpower/description/
// 给你一个下标从 0 开始的二维整数数组 questions ，其中 questions[i] = [pointsi, brainpoweri] 。
//
// 这个数组表示一场考试里的一系列题目，你需要 按顺序 （也就是从问题 0 开始依次解决），针对每个问题选择 解决 或者 跳过 操作。解决问题 i 将让你 获得  pointsi 的分数，但是你将 无法 解决接下来的 brainpoweri 个问题（即只能跳过接下来的 brainpoweri 个问题）。如果你跳过问题 i ，你可以对下一个问题决定使用哪种操作。
//
// 比方说，给你 questions = [[3, 2], [4, 3], [4, 4], [2, 5]] ：
// 如果问题 0 被解决了， 那么你可以获得 3 分，但你不能解决问题 1 和 2 。
// 如果你跳过问题 0 ，且解决问题 1 ，你将获得 4 分但是不能解决问题 2 和 3 。
// 请你返回这场考试里你能获得的 最高 分数。

// 时间复杂度太高
//func mostPoints(questions [][]int) int64 {
//	f := make([]int64, len(questions)+1)
//	for i := 1; i < len(questions)+1; i++ {
//		solved := int64(questions[i-1][0])
//		for j := i - 2; j >= 0; j-- {
//			if i-1-j > questions[j][1] {
//				solved = max(solved, int64(questions[i-1][0])+f[j+1])
//			}
//		}
//		f[i] = max(f[i], solved)
//	}
//	return slices.Max(f)
//}

// 倒序DP
func mostPoints(questions [][]int) int64 {
	f := make([]int64, len(questions)+1)
	for i := len(questions) - 1; i >= 0; i-- {
		f[i] = Max(f[i+1], int64(questions[i][0])+f[min(len(questions), i+questions[i][1]+1)])
	}
	return f[0]
}

func TestMostPoints() {
	questions := [][]int{
		{12, 46},
		{78, 19},
		{63, 15},
		{79, 62},
		{13, 10},
	}
	fmt.Printf("Input: questions=%v\n", questions)
	fmt.Println("你能获得的 最高 分数: ", mostPoints(questions))
}
