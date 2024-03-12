package dynamic_programming

import (
	"encoding/json"
	"fmt"
)

// ==== some define ====

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

// ==== some helper function ====
func printArr2(arr [][]int) {
	for _, line := range arr {
		fmt.Println(line)
	}
}

func printArr3(arr [][][]int) {
	for _, arr2 := range arr {
		printArr2(arr2)
	}
}

func tree2JSON(root *TreeNode) (string, error) {
	jsonBytes, err := json.Marshal(root)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
func trees2JSON(tres []*TreeNode) (string, error) {
	jsonBytes, err := json.Marshal(tres)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
