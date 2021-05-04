package main

import "fmt"

func sort(arr []int) []int {
	// base
	if len(arr) == 1 {
		return arr
	}
	// hypothesis
	temp := arr[len(arr)-1]
	arr = sort(arr[0:len(arr)-1])
	// induction
	return insert(arr, temp)
}

func insert(arr []int, temp int) []int {
	// base
	if len(arr) == 0 || arr[len(arr)-1] <= temp {
		arr = append(arr, temp)
		return arr
	}
	// hypothesis
	val := arr[len(arr)-1]
	insert(arr[0:len(arr)-1], temp)
	// induction
	return append(arr, val)
}

func main() {
	a := []int{6,7,1,2,9,5,6}
	fmt.Println(sort(a))
}