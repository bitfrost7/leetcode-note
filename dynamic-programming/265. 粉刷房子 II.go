package dynamic_programming

import (
	"fmt"
	"math"
	"slices"
)

//https://leetcode.cn/problems/paint-house-ii/description/
//假如有一排房子共有 n 幢，每个房子可以被粉刷成 k 种颜色中的一种。房子粉刷成不同颜色的花费成本也是不同的。你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
//
//每个房子粉刷成不同颜色的花费以一个 n x k 的矩阵表示。
//
//例如，costs[0][0] 表示第 0 幢房子粉刷成 0 号颜色的成本；costs[1][2] 表示第 1 幢房子粉刷成 2 号颜色的成本，以此类推。
//返回 粉刷完所有房子的最低成本 。

func minCostII(costs [][]int) int {
	n, k := len(costs), len(costs[0])
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	for i := 1; i < n+1; i++ {
		for j := 0; j < k; j++ {
			f[i][j] = costs[i-1][j] + slices.Min(slices.Concat(f[i-1][:j], f[i-1][j+1:]))
		}
	}
	return slices.Min(f[len(costs)])
}

func minNoIndex(arr []int, k int) int {
	minn := math.MaxInt
	for i := range arr {
		if i != k {
			minn = min(minn, arr[i])
		}
	}
	return minn
}

func TestMinCostII() {
	costs := [][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}
	//costs := [][]int{
	//	{1, 3},
	//	{2, 4},
	//}
	fmt.Printf("Input: costs=%v\n", costs)
	fmt.Println("粉刷完所有房子最少的花费成本", minCostII(costs))
}
