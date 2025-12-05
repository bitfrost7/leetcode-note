package helpers

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

func PrintTree(root *TreeNode) {
	printTree(root, "", true)
}

func PrintTrees(trees []*TreeNode) {
	for _, node := range trees {
		PrintTree(node)
	}
}

func printTree(node *TreeNode, prefix string, isRight bool) {
	if node == nil {
		return
	}

	if prefix != "" {
		fmt.Printf("%s", prefix)
		if isRight {
			fmt.Print("└── ")
		} else {
			fmt.Print("├── ")
		}
	}
	fmt.Println(node.Val)

	newPrefix := prefix
	if prefix != "" {
		if isRight {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
	}

	printTree(node.Right, newPrefix, false)
	printTree(node.Left, newPrefix, true)
}


func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}