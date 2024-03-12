package math

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32; i++ {
		bit := num & 1 //num的最低位
		rev = rev | bit<<(31-i)
		num = num >> 1
	}
	return rev
}

// @lc code=end
