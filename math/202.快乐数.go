package math

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
//详见https://leetcode.cn/problems/happy-number/solutions/224894/kuai-le-shu-by-leetcode-solution/
// 任一n的平方和 要么最终等于1 要么在一个循环里

func isHappy(n int) bool {
	step := func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return sum
	}
	m := map[int]bool{}
	for n != 1 {
		n = step(n)
		if m[n] {
			return false
		} else {
			m[n] = true
		}
	}
	return n == 1

}

// @lc code=end
