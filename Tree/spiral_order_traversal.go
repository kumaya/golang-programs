package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

func Height(node *Node) int {
	if node == nil {
		return 0
	}
	leftHeight := Height(node.left)
	rightHeight := Height(node.right)
	if leftHeight >= rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}

func printSpiral(root *Node, level int, flag bool) {
	if root == nil {
		return
	}
	if level == 0 {
		fmt.Print(" ", root.data)
	}
	if flag {
		printSpiral(root.left, level-1, flag)
		printSpiral(root.right, level-1, flag)
	} else {
		printSpiral(root.right, level-1, flag)
		printSpiral(root.left, level-1, flag)
	}
}

func SpiralView(root *Node, clockwise bool) {
	if root == nil {
		return
	}
	for i:=0; i<Height(root);i++ {
		printSpiral(root, i, clockwise)
		clockwise = !clockwise
	}
}

/*
    1
  2   3
 4     5
        6
*/
func main() {
	root := Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.right.right = &Node{data: 5}
	root.right.right.right = &Node{data: 6}
	fmt.Println("Clockwise spiral tree:")
	SpiralView(&root, true)
	fmt.Println()
	fmt.Println("Anti-Clockwise spiral tree:")
	SpiralView(&root, false)
}