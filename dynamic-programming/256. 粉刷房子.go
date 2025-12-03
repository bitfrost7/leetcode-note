package dynamic_programming

import (
	"fmt"
	. "leetcode-note/helpers"
)

//https://leetcode.cn/problems/paint-house/description/
// 假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
//
//当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。
//
//例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。
//
//请计算出粉刷完所有房子最少的花费成本。

func minCost(costs [][]int) int {
	f := make([][3]int, len(costs)+1)
	for i := 1; i < len(costs)+1; i++ {
		f[i][0] = costs[i-1][0] + min(f[i-1][1], f[i-1][2])
		f[i][1] = costs[i-1][1] + min(f[i-1][0], f[i-1][2])
		f[i][2] = costs[i-1][2] + min(f[i-1][0], f[i-1][1])
	}
	minn := 21 * len(costs)
	for _, cost := range f[len(costs)] {
		minn = min(minn, cost)
	}
	return minn
}

func minCost1(costs [][]int) int {
	r, b, g := 0, 0, 0
	for i := 1; i < len(costs)+1; i++ {
		r, b, g = costs[i-1][0]+min(b, g), costs[i-1][1]+min(r, g), costs[i-1][2]+min(r, b)
	}
	return Min3(r, b, g)
}

func TestMinCost() {
	costs := [][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}
	//costs := [][]int{
	//	{7, 6, 2},
	//}
	fmt.Printf("Input: costs=%v\n", costs)
	fmt.Println("粉刷完所有房子最少的花费成本", minCost(costs))
}
