package main

import (
	"fmt"
	"math"
)

type Node struct {
	left  *Node
	right *Node
	data  int
}

func HeightOfNode(node *Node, d int) int {
	dist := 0
	var q []*Node
	q = append(q, node)
	q = append(q, &Node{data: math.MaxInt32})
	for len(q) > 0 {
		this := q[0]
		q = q[1:]
		if this == nil {
			continue
		}
		if this.data == math.MaxInt32 && len(q) > 1 {
			q = append(q, &Node{data: math.MaxInt32})
			dist += 1
		} else {
			if this.data == d {
				return dist
			}
			q = append(q, this.left)
			q = append(q, this.right)
		}
	}
	return -1
}

func printAtHeight(root *Node, this *Node, h int, cheight int) {
	if root == nil {
		return
	}
	if h == cheight+1 {
		if root.left != nil && root.left != this {
			fmt.Print("==", root.left.data)
		}
		if root.right != nil && root.right != this {
			fmt.Print("==", root.right.data)
		}
	}
	printAtHeight(root.left, this, h, cheight+1)
	printAtHeight(root.right, this, h, cheight+1)
}

/*
      1
   2    3
 4  7    5
          6
*/
func main() {
	root := Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 7}
	root.right.right = &Node{data: 5}
	root.right.right.right = &Node{data: 6}
	n := root.left.left
	fmt.Printf("Cousins of the node %d:\n", n.data)
	h := HeightOfNode(&root, n.data)
	printAtHeight(&root, n, h, 0)
}
