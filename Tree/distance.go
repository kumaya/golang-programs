package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

func distanceFromRoot(root *Node, data int) int {
	if root == nil {
		return -1
	}
	dist := -1
	if root.data == data {
		return 0
	}
	ld := distanceFromRoot(root.left, data)
	if ld >= 0 {
		return 1 + ld
	}
	rd := distanceFromRoot(root.right, data)
	if rd >= 0 {
		return 1 + rd
	}
	return dist
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
	fmt.Println("Distance from root:")
	fmt.Println(
		distanceFromRoot(&root, 1),
		distanceFromRoot(&root, 2),
		distanceFromRoot(&root, 3),
		distanceFromRoot(&root, 4),
		distanceFromRoot(&root, 5),
		distanceFromRoot(&root, 6),
		distanceFromRoot(&root, 7))
}
