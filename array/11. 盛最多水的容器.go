package array

import "fmt"

// description: (https://leetcode.cn/problems//description/)
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 返回容器可以储存的最大水量。

// 说明：你不能倾斜容器。

// ----- Code Here -----

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxWater := 0
	for l < r {
		hL, hR := height[l], height[r]
		minH := min(hL, hR)
		area := (r - l) * minH
		if area > maxWater {
			maxWater = area
		}

		// 移动较短的那一边
		if hL <= hR {
			l++
		} else {
			r--
		}
	}
	return maxWater
}

// ----- Begin Test -----

func TestMaxArea() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("height", height)
	fmt.Println("maxArea:", maxArea(height)) // 49
}

// ----- Solution Here -----
/*
核心思路: 梯度最缓下降法+双指针

使用双指针法，从数组的两端开始，计算当前容器的面积，并记录最大面积。然后移动较短的那一边的指针，尝试找到更大的面积。重复这个过程直到两个指针相遇。
移动双指针的过程就是面积下降的过程，我们需要让整个下降过程最慢，因为面积的变化取决于较短边的高度，移动较短边有可能找到更高的边，从而增加面积。

*/
