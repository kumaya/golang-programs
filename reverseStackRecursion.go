package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) push(item int) {
	*s = append(*s, item)
}

func (s *Stack) pop() int {
	idx := len(*s) -1
	element := (*s)[idx]
	*s = (*s)[:idx]
	return element
}

func reverseStack(stk Stack) Stack {
	// base
	if len(stk) == 0{
		return stk
	}
	// hypothesis
	temp := stk.pop()
	stk = reverseStack(stk)
	// induction
	stk = insert(stk, temp)
	return stk
}

func insert(stk Stack, item int) Stack {
	// base
	if len(stk) == 0 {
		stk.push(item)
		return stk
	}
	// hypothesis
	val := stk.pop()
	stk = insert(stk, item)
	// induction
	stk.push(val)
	return stk
}

func main() {
	var stack Stack
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(5)
	stack.push(4)
	fmt.Println("Input stack: ", stack)
	fmt.Println("Reversed stack: ", reverseStack(stack))
}
