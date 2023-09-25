package dynamic_programming

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min3(a, b, c int) int {
	_min := min(a, b)
	return min(_min, c)
}
