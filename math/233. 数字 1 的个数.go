package math

// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
//
// 示例 1：
//
// 输入：n = 13
// 输出：6
// 示例 2：
//
// 输入：n = 0
// 输出：0
//
// 提示：
//
// 0 <= n <= 109
func countDigitOne(n int) (ans int) {
	for k, mulk := 0, 1; n >= mulk; k++ {
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
		mulk *= 10
	}
	return
}
