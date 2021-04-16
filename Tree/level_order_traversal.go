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

func recursiveLevelOrder(root *Node, level int) {
	if root == nil {
		return
	}
	if level == 0 {
		fmt.Print(" ", root.data)
	}
	recursiveLevelOrder(root.left, level-1)
	recursiveLevelOrder(root.right, level-1)
}

func iterativeLevelOrder(root *Node) {
	if root == nil {
		return
	}
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		this := q[0]
		q = q[1:]
		if this.left != nil {
			q = append(q, this.left)
		}
		if this.right != nil {
			q = append(q, this.right)
		}
		fmt.Print(" ", this.data)
	}
}

/*
    1
  2   3
 4     5
      6 7
*/
func main() {
	root := Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.right.right = &Node{data: 5}
	root.right.right.left = &Node{data: 6}
	root.right.right.right = &Node{data: 7}

	fmt.Println("Recursive level order traversal:")
	for i := 0; i < Height(&root); i++ {
		recursiveLevelOrder(&root, i)
	}
	fmt.Println()
	fmt.Println("Iterative level order traversal:")
	iterativeLevelOrder(&root)
}
