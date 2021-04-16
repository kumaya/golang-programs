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

func recursiveReverseLevelOrder(root *Node, level int) {
	if root == nil {
		return
	}
	if level == 0 {
		fmt.Print(" ", root.data)
	}
	recursiveReverseLevelOrder(root.left, level-1)
	recursiveReverseLevelOrder(root.right, level-1)
}

func iterativeReverseLevelOrder(root *Node) {
	if root == nil {
		return
	}
	var q, stack []*Node
	q = append(q, root)
	for len(q) > 0 {
		this := q[0]
		stack = append(stack, this)
		q = q[1:]
		if this.right != nil {
			q = append(q, this.right)
		}
		if this.left != nil {
			q = append(q, this.left)
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Print(" ", stack[i].data)
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
	fmt.Println("Recursive reverse level order traversal:")
	for i := Height(&root); i >= 0; i-- {
		recursiveReverseLevelOrder(&root, i)
	}

	fmt.Println()
	fmt.Println("Iterative reverse level order traversal:")
	iterativeReverseLevelOrder(&root)
}
