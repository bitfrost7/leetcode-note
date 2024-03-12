package math

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	ones := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ones++
		}
	}
	return ones
}

// @lc code=end
