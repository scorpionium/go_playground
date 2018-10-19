package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

func AddElement(root *Node, v int) *Node {
	node := Node{Val: v}
	if root == nil {
		return &node
	} else {
		if v <= root.Val {
			cur := AddElement(root.Left, v)
			root.Left = cur
		} else {
			cur := AddElement(root.Right, v)
			root.Right = cur
		}
		return root
	}
}

func Height(root *Node) int {
	if root == nil {
		return 0
	} else {
		hl := Height(root.Left)
		hr := Height(root.Right)
		if hl > hr {
			return hl + 1
		} else {
			return hr + 1
		}
	}
}

func InOrderTraversalPrint(root *Node) {
	if root != nil {
		InOrderTraversalPrint(root.Left)
		fmt.Printf("%d ", root.Val)
		InOrderTraversalPrint(root.Right)
	}
}

func LevelTraversalPrint(root *Node) {
	h := Height(root)
	for i := 1; i <= h; i++ {
		PrintLevel(root, i)
		fmt.Println()
	}
}

func PrintLevel(root *Node, l int) {
	if root != nil {
		if l == 1 {
			fmt.Printf("%d ", root.Val)
		} else {
			PrintLevel(root.Left, l-1)
			PrintLevel(root.Right, l-1)
		}
	}
}

func main() {
	input := []int{25, 20, 11, 30, 17, 35, 14, 40, 15}
	var root *Node
	for _, v := range input {
		root = AddElement(root, v)
	}
	fmt.Printf("Height: %d\n", Height(root))
	InOrderTraversalPrint(root)
	fmt.Println()
	LevelTraversalPrint(root)
}
