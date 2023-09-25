package dynamic_programming

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	arr := make([]int, max+1)
	for _, v := range nums {
		arr[v] += v
	}
	return rob(arr)
}
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	curr, prev := max(nums[0], nums[1]), nums[0]
	for i := 2; i < len(nums); i++ {
		curr, prev = max(curr, prev+nums[i]), curr
	}
	return curr
}
