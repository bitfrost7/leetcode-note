package digit

import "fmt"

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i, num := range nums {
		if bitCount(i) == k {
			sum += num
		}
	}
	return sum
}

func bitCount(x int) int {
	count := 0
	for x != 0 {
		if x%2 == 1 {
			count++
		}
		x /= 2
	}
	return count
}

func TestSumIndicesWithKSetBits() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	fmt.Println(nums)
	fmt.Println(bitCount(5)) // 101
	fmt.Println(sumIndicesWithKSetBits(nums, k))
}
