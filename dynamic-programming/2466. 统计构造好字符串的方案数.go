package dynamic_programming

// 给你整数 zero ，one ，low 和 high ，我们从空字符串开始构造一个字符串，每一步执行下面操作中的一种：
//
// 将 '0' 在字符串末尾添加 zero  次。
// 将 '1' 在字符串末尾添加 one 次。
// 以上操作可以执行任意次。
//
// 如果通过以上过程得到一个 长度 在 low 和 high 之间（包含上下边界）的字符串，那么这个字符串我们称为 好 字符串。
//
// 请你返回满足以上要求的 不同 好字符串数目。由于答案可能很大，请将结果对 109 + 7 取余 后返回。
func countGoodStrings(low int, high int, zero int, one int) int {
	const MOD = int(1e9 + 7)
	f := make([]int, high+1)
	f[0] = 1
	ans := 0
	for i := 1; i < high+1; i++ {
		if i >= zero {
			f[i] = (f[i] + f[i-zero]) % MOD
		}
		if i >= one {
			f[i] = (f[i] + f[i-one]) % MOD
		}
		if i >= low {
			ans = (ans + f[i]) % MOD
		}
	}
	return ans
}
