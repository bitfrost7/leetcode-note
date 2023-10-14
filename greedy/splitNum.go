package greedy

import (
	"fmt"
	"sort"
)

/*
给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：

num1 和 num2 直接连起来，得到 num 各数位的一个排列。
换句话说，num1 和 num2 中所有数字出现的次数之和等于 num 中所有数字出现的次数。
num1 和 num2 可以包含前导 0 。
请你返回 num1 和 num2 可以得到的和的 最小 值。

注意：
num 保证没有前导 0 。
num1 和 num2 中数位顺序可以与 num 中数位顺序不同。
*/
func splitNum(num int) int {
	if num < 10 {
		return num
	}
	cnt := make([]int, 9)
	n := 0
	for num > 0 {
		cnt[num%10-1]++
		num /= 10
		n++
	}
	sort.Ints(cnt)
	ans := [2]int{}
	for i, j := 0, 0; i < n; i++ {
		for cnt[j] == 0 {
			j++
		}
		cnt[j]--
		ans[i&1] = ans[i&1]*10 + j
	}
	return ans[0] + ans[1]
}
func TestSplitNum() {
	fmt.Println(splitNum(687))
}
