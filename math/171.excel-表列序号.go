package math

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	res := 0
	for i := 0; i < len(columnTitle); i++ {
		res = res*26 + int(columnTitle[i]-'A'+1)
	}
	return res
}

// @lc code=end
