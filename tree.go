package main

import (
	"fmt"
    "math"
    "sort"
)

type Node struct {
	data  int
	left  *Node
	right *Node
	hDist int
}

func height(root *Node) float64 {
	if root == nil {
		return 0
	}
	l := height(root.left)
	r := height(root.right)
	return math.Max(l, r) + 1.0
}

func diameter(root *Node) float64 {
    if root == nil {
        return 0
    }
    lH := height(root.left)
    rH := height(root.right)
    lD := diameter(root.left)
    rD := diameter(root.right)
    maxD := math.Max(lD, rD)
    return math.Max(maxD, lH+rH+1)
}

func isLeaf(root *Node) bool {
	if root == nil {
		return false
	}
	if root.left == nil && root.right == nil {
		return true
	}
	return false
}

func sum(root *Node) int {
	if root == nil {
		return 0
	}
	return sum(root.left) + root.data + sum(root.right)
}

func sumTree(root *Node) bool {
	if root == nil {
		return true
	}
	if isLeaf(root) {
		return true
	}
	lsum := sum(root.left)
	rsum := sum(root.right)
	if (lsum+rsum == root.data) && sumTree(root.left) && sumTree(root.right) {
		return true
	}
	return false
}

func lca(root *Node, x, y int) *Node {
	if root == nil {
		return nil
	}
	if root.data == x || root.data == y {
		return root
	}
	l := lca(root.left, x, y)
	r := lca(root.right, x, y)
	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	} else {
		return r
	}
}

func topView(root *Node) {
	if root == nil {
		return
	}
	hd := make(map[int]int)
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		var n *Node
		n, q = q[0], q[1:]
		if _, ok := hd[n.hDist]; !ok {
			hd[n.hDist] = n.data
		}
		//hd[n.hDist] = n.data
		if n.left != nil {
			n.left.hDist = n.hDist - 1
			q = append(q, n.left)
		}
		if n.right != nil {
			n.right.hDist = n.hDist + 1
			q = append(q, n.right)
		}
	}
	var keys []int
	for k := range hd {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		fmt.Printf("%d ", hd[v])
	}
}

func isMirror(r1, r2 *Node) bool {
	if r1 == nil || r2 == nil {
		return true
	}
	if r1 != nil && r2 != nil {
		if r1.data == r2.data {
			return isMirror(r1.left, r2.left) && isMirror(r1.right, r2.right)
		}
	}
	return false
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.data)
	inorder(root.right)
}

func inorderIter(root *Node) {
	var stack []*Node
	stack = append(stack, root)
	for len(stack) > 0{
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(top.data)
		stack = append(stack, top.right)
		}
	}

func main() {
	root := Node{26,
		&Node{10,
			&Node{4, nil, nil, 0},
			&Node{6, nil, nil, 0}, 0},
		&Node{3,
			nil,
			&Node{3, nil, nil, 0}, 0},
		0,
	}
	fmt.Printf("Is sum tree: %t\n", sumTree(&root))
	x := 6
	y := 3
	fmt.Printf("lca of %d and %d = %d\n", x, y, lca(&root, x, y).data)

	fmt.Printf("top view of a tree: ")
	topView(&root)

	root1 := root
	//root1.data = 4
	fmt.Printf("\nisMirror: %t", isMirror(&root, &root1))

	fmt.Printf("\nDiameter: %f", diameter(&root))

	fmt.Printf("Inorder recursive: \n")
	inorder(&root)
	fmt.Println("")
	inorderIter(&root)
}
