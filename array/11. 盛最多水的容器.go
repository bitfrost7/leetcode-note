package array

import (
	"fmt"
	. "leetcode-note/helpers"
)

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
		minH := Min(hL, hR)
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

// 解法：双指针
//
// 容器面积 = (r-l) * min(height[l], height[r])。
// 初始化左右指针 l=0, r=n-1，计算当前面积并更新最大值。
//
// 指针移动策略：
// - 总是移动“较短的那一侧”。因为面积受短板限制，移动长板不可能让 min 变大，宽度还会变小。
// - 移动短板才有机会找到更高的短板，从而提升面积。
//
// 复杂度：
// - 时间 O(n)
// - 额外空间 O(1)
