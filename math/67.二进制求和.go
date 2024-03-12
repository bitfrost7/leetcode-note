/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
package math

// @lc code=start
func addBinary(a string, b string) string {
	var ans string
	var carry int
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		var sum int
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		sum += carry
		ans = string(sum%2+'0') + ans
		carry = sum / 2
	}
	return ans
}

// @lc code=end
