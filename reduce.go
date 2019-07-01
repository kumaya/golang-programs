package main

import (
	"fmt"
)

type reducer func(a, b int) int

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func reduce(lst []int, f reducer) int {
	res := lst[0]
	for i := 1; i < len(lst); i++ {
		res = f(res, lst[i])
	}
	return res
}

func main() {
	lst := []int{5, 1, 2, 10}
	fmt.Println(reduce(lst, add))
	fmt.Println(reduce(lst, multiply))
}
