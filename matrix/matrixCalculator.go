package main

import "fmt"
import "matrix/operations"


func main() {
	rawFirstMatrix := matrix.GenerateMatrix(3, 4)
	rawSecondMatrix := matrix.GenerateMatrix(3, 4)
	firstMatrix := matrix.PopulateMatrix(rawFirstMatrix)
	secondMatrix := matrix.PopulateMatrix(rawSecondMatrix)
	fmt.Println("First matrix: ", firstMatrix)
	fmt.Println("Second matrix: ", secondMatrix)
	fmt.Println("")
	
	fmt.Println("*****ADDITION*****")
	addRes, isValid := matrix.Add(firstMatrix, secondMatrix)
	fmt.Println("Matrix Addition is valid: ", isValid)
	fmt.Println("Matrix Addition Result: ", addRes)
	fmt.Println("")
	
	fmt.Println("*****SUBTRACTION*****")
	subtractRes, isValid := matrix.Subtract(firstMatrix, secondMatrix)
	fmt.Println("Matrix Subtraction is valid: ", isValid)
	fmt.Println("Matrix Subtraction Result: ", subtractRes)
	fmt.Println("")

	fmt.Println("*****MULTIPLICATION*****")
	rawFirstMatrix = matrix.GenerateMatrix(3, 4)
	rawSecondMatrix = matrix.GenerateMatrix(3, 4)
	firstMatrix = matrix.PopulateMatrix(rawFirstMatrix)
	secondMatrix = matrix.PopulateMatrix(rawSecondMatrix)
	multiplyRes, isValid := matrix.Multiply(firstMatrix, secondMatrix)
	fmt.Println("First matrix: ", firstMatrix)
	fmt.Println("Second matrix: ", secondMatrix)
	fmt.Println("")
	fmt.Println("Matrix Multiplication is valid: ", isValid)
	fmt.Println("Matrix Multiplication Result: ", multiplyRes)
	fmt.Println("")
}