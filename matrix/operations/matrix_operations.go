package matrix

import "fmt"

// Matrix is a slice of slice of int of a specified size
type Matrix [][]int

// New creates a new matrix with r rows and c columns
func New(r int, c int) [][]int {
	mat := make([][]int, r)
	for i := range mat {
		mat[i] = make([]int, c)
	}
	return mat
}

/*
func PopulateMatrix(m [][]int) ([][]int){
	len_r := len(m)
	for i:=0; i<len_r; i++ {
		c := m[i]
		len_c := len(c)
		for j:=0; j<len_c; j++ {
			m[i][j] = rand.Intn(3)
		}
	}
	return m
}
*/

func isvalid(m1 [][]int, m2 [][]int) (valid bool) {
	m1_len_r, m1_len_c, m2_len_r, m2_len_c := len(m1), len(m1[0]), len(m2), len(m2[0])
	valid = (m1_len_r == m2_len_r && m1_len_c == m2_len_c)
	return
}

func iterate(m1 [][]int, operation func(int, int)) {
	for i, m1row := range m1 {
		for j := range m1row {
			operation(i, j)
		}
	}
}

// Add adds two matrices if their dimensions are the same
// It returns a new matrix with the addition results,
//   and the validity of the operation
func Add(m1 [][]int, m2 [][]int) (res [][]int, valid bool) {
	valid = isvalid(m1, m2)
	if valid {
		res = New(len(m1), len(m1[0]))

		iterate(m1, func(i int, j int) {
			res[i][j] = m1[i][j] + m2[i][j]
		})
	}

	return
}

// Subtract subtracts two matrices, m2 from m1,  if their dimensions are the same
// It returns a new matrix with the addition results,
//   and the validity of the operation
func Subtract(m1 [][]int, m2 [][]int) (res [][]int, valid bool) {
	valid = isvalid(m1, m2)
	if valid {
		res = New(len(m1), len(m1[0]))

		iterate(m1, func(i int, j int) {
			res[i][j] = m1[i][j] - m2[i][j]
		})
	}

	return
}

// String returns a string representation of the matrix
func String(m [][]int) string {
	result := ""

	iterate(m, func(i int, j int) {
		delim := "\t"
		if j == 0 {
			delim = "\n"
		}

		result = result + fmt.Sprintf("%s%4d", delim, m[i][j])

	})

	result += "\n"

	return result
}

/*
func Multiply(m1 [][]int, m2 [][]int) ([][]int, bool){
	m1_len_r := len(m1)
	m1_len_c := len(m1[0])
	m2_len_r := len(m2)
	m2_len_c := len(m2[0])
	res := GenerateMatrix(m1_len_r, m2_len_c)
	var valid bool
	valid = true
	// fmt.Println(m1_len_r, m2_len_r, m1_len_c, m2_len_c)
	if m1_len_c != m2_len_r {
		valid = false
	}
	if valid {
		var temp int
		temp = 0
		for i:=0; i<m1_len_r; i++ {
			for j:=0; j<m1_len_c; j++ {
				for k:=0; k<m1_len_c; k++ {
					temp += m1[i][k] * m2[k][j]
				}
				res[i][j] = temp
				temp = 0
			}
		}
	}
	return res, valid
}
*/
