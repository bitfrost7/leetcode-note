package dynamic_programming

import "fmt"

// 给你一个整数 n，请你帮忙统计一下我们可以按下述规则形成多少个长度为 n 的字符串：
//
// 字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
// 每个元音 'a' 后面都只能跟着 'e'
// 每个元音 'e' 后面只能跟着 'a' 或者是 'i'
// 每个元音 'i' 后面 不能 再跟着另一个 'i'
// 每个元音 'o' 后面只能跟着 'i' 或者是 'u'
// 每个元音 'u' 后面只能跟着 'a'
// 由于答案可能会很大，所以请你返回 模 10^9 + 7 之后的结果。

func countVowelPermutation(n int) int {
	const MOD int = 1e9 + 7
	f := make([][5]int, n+1)
	f[n][0] = 1 // 'a'
	f[n][1] = 1 // 'e'
	f[n][2] = 1 // 'i'
	f[n][3] = 1 // 'o'
	f[n][4] = 1 // 'u'
	for i := n - 1; i >= 0; i-- {
		f[i][0] = f[i+1][1] % MOD
		f[i][1] = (f[i+1][0] + f[i+1][2]) % MOD
		f[i][2] = (f[i+1][0] + f[i+1][1] + f[i+1][3] + f[i+1][4]) % MOD
		f[i][3] = (f[i+1][2] + f[i+1][4]) % MOD
		f[i][4] = f[i+1][0] % MOD
	}
	sum := 0
	for _, num := range f[1] {
		sum = (sum + num) % MOD
	}
	return sum
}

func countVowelPermutation1(n int) int {
	const MOD int = 1e9 + 7
	a, e, i, o, u := 1, 1, 1, 1, 1
	for k := 1; k < n; k++ {
		a, e, i, o, u = e, (a+i)%MOD, (a+e+o+u)%MOD, (i+u)%MOD, a
	}
	return (a + e + i + o + u) % MOD
}
func TestCountVowelPermutation() {
	n := 144
	fmt.Printf("Input: n=%v\n", n)
	fmt.Println("粉刷完所有房子最少的花费成本", countVowelPermutation1(n))
}
