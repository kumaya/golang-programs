package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

func LeftView(root *Node, currentLevel int, levelTraversed []int) {
	if root == nil {
		return
	}
	if currentLevel > levelTraversed[0] {
		fmt.Print(" ", root.data)
		levelTraversed[0] = currentLevel
	}
	LeftView(root.left, currentLevel+1, levelTraversed)
	LeftView(root.right, currentLevel+1, levelTraversed)
}

func RightView(root *Node, currentLevel int, levelTraversed []int) {
	if root == nil {
		return
	}
	if currentLevel > levelTraversed[0] {
		fmt.Print(" ", root.data)
		levelTraversed[0] = currentLevel
	}
	RightView(root.right, currentLevel+1, levelTraversed)
	RightView(root.left, currentLevel+1, levelTraversed)
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
	fmt.Println("Left view of a tree:")
	var i = make([]int, 1, 1)
	i[0] = -1
	LeftView(&root, 0, i)

	fmt.Println()
	fmt.Println("Right view of a tree:")
	var j = make([]int, 1, 1)
	j[0] = -1
	RightView(&root, 0, j)
}
