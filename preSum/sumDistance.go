package preSum

import (
	"fmt"
	"sort"
)

/*
有一些机器人分布在一条无限长的数轴上，他们初始坐标用一个下标从 0 开始的整数数组 nums 表示。当你给机器人下达命令时，它们以每秒钟一单位的速度开始移动。

给你一个字符串 s ，每个字符按顺序分别表示每个机器人移动的方向。'L' 表示机器人往左或者数轴的负方向移动，'R' 表示机器人往右或者数轴的正方向移动。

当两个机器人相撞时，它们开始沿着原本相反的方向移动。

请你返回指令重复执行 d 秒后，所有机器人之间两两距离之和。由于答案可能很大，请你将答案对 109 + 7 取余后返回。
*/
func sumDistance(nums []int, s string, d int) int {
	const mod int = 1e9 + 7
	n := len(nums)
	pos := make([]int, n)
	for i, ch := range s {
		if ch == 'L' {
			pos[i] = nums[i] - d
		} else {
			pos[i] = nums[i] + d
		}
	}
	sort.Ints(pos)
	res := 0
	for i := 1; i < n; i++ {
		res += (pos[i] - pos[i-1]) * i % mod * (n - i) % mod // 两两距离之和
		/*
			假设位置为[a,b,c,d,e],要求两两距离之和

			a到其他位置的距离之和
			a + b
			a + b + c
			a + b + c + d
			a + b + c + d + e

			b到其他位置距离之和
			    b + c
			    b + c + d
			    b + c + d + e

			c到其他位置距离之和
					c + d
					c + d + e

			d到其他位置距离之和
						d + e

			可以发现规律：
			两两距离之和 =
			(a+b) * 1 * (n - 1) +
			(b+c) * 2 * (n - 2) +
			(c+d) * 3 * (n - 3) +
			(d+e) * 3 * (n - 3)
		*/
		res %= mod
	}
	return res
}

/*
前缀和: 快速计算区间和的一种技巧
假设arr是一个长度为n的数组 arr = 1...n
preSum 是一个长度为n+1的数组, 为了避免讨论i=0的情况, 令len(preSum) = n+1, 可以用preSum[1] = preSum[0] + arr[0]来表示前缀和,一般模板如下:
		for i := 1; i <= n; i++ {
			preSum[i] = preSum[i-1] + arr[i-1]
		}
基于前缀和的区间和计算:
区间和 sum(i, j) = preSum[j+1] - preSum[i]
preSum[j+1] 0到j   的和 包括j
preSum[i] 	0到i-1 的和 不包括i
preSum[j+1] - preSum[i] i到j 的和
*/

func TestSumDistance() {
	fmt.Println(sumDistance([]int{1, 2, 3, 4, 5}, "RRRLL", 3))
}
