package main

import "fmt"
import "github.com/kumaya/golang-programs/matrix/operations"
import "math/rand"

/*
func populateMatrix(m [][]int) matrix.Matrix {
	len_r := len(m)
	for i := 0; i < len_r; i++ {
		c := m[i]
		len_c := len(c)
		for j := 0; j < len_c; j++ {
			m[i][j] = rand.Intn(3)
		}
	}
	return m
}
*/

func populateMatrix(m matrix.Matrix) matrix.Matrix {
	rows, cols := m.Size()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Set(i, j, rand.Intn(3))
		}
	}
	return m
}

func main() {
	rawFirstMatrix := matrix.New(3, 3)
	rawSecondMatrix := matrix.New(3, 3)
	firstMatrix := populateMatrix(rawFirstMatrix)
	secondMatrix := populateMatrix(rawSecondMatrix)
	fmt.Println("First matrix: ", firstMatrix)
	fmt.Println("Second matrix: ", secondMatrix)
	fmt.Println("")

	fmt.Println("*****ADDITION*****")
	addRes, isValid := matrix.Add(firstMatrix, secondMatrix)
	fmt.Println("Matrix Addition is valid: ", isValid)
	fmt.Println("Matrix Addition Result: ", matrix.String(addRes))
	fmt.Println("")

	fmt.Println("*****SUBTRACTION*****")
	subtractRes, isValid := matrix.Subtract(firstMatrix, secondMatrix)
	fmt.Println("Matrix Subtraction is valid: ", isValid)
	fmt.Println("Matrix Subtraction Result: ", matrix.String(subtractRes))
	fmt.Println("")

	firstMatrix.Set(2, 2, 67)
	fmt.Println(firstMatrix)

	firstMatrix.Add(secondMatrix)
	fmt.Println(firstMatrix)

	/*
		fmt.Println("*****MULTIPLICATION*****")
		rawFirstMatrix = matrix.GenerateMatrix(3, 4)
		rawSecondMatrix = matrix.GenerateMatrix(3, 4)
		firstMatrix = populateMatrix(rawFirstMatrix)
		secondMatrix = populateMatrix(rawSecondMatrix)
		multiplyRes, isValid := matrix.Multiply(firstMatrix, secondMatrix)
		fmt.Println("First matrix: ", firstMatrix)
		fmt.Println("Second matrix: ", secondMatrix)
		fmt.Println("")
		fmt.Println("Matrix Multiplication is valid: ", isValid)
		fmt.Println("Matrix Multiplication Result: ", multiplyRes)
		fmt.Println("")
	*/
}
