// Package helpers 提供通用的辅助函数
package helpers

import (
	"cmp"
	"fmt"
)

// PrintAnyArr 打印任意类型的一维数组
func PrintAnyArr[T any](arr []T) {
	fmt.Print("[")
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

// PrintArr2 打印任意类型的二维数组
func PrintArr2[T any](arr [][]T) {
	fmt.Println("[")
	for _, row := range arr {
		fmt.Print("  ")
		PrintAnyArr(row)
	}
	fmt.Println("]")
}

// PrintArr3 打印任意类型的三维数组
func PrintArr3[T any](arr [][][]T) {
	fmt.Println("[")
	for _, matrix := range arr {
		fmt.Print("  ")
		PrintArr2(matrix)
	}
	fmt.Println("]")
}

func Max[K cmp.Ordered](a, b K) K {
	if a > b {
		return a
	}
	return b
}

func Min[K cmp.Ordered](a, b K) K {
	if a < b {
		return a
	}
	return b
}

func Max3[K cmp.Ordered](a, b, c K) K {
	return Max(Max(a, b), c)
}

func Min3[K cmp.Ordered](a, b, c K) K {
	return Min(Min(a, b), c)
}
