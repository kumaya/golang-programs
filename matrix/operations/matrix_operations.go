// matrix is a package that defines a matrix type, and operations on it.
package matrix

import "fmt"

// Matrix is a slice of slice of int of a specified size
type Matrix struct {
	data [][]int
}

// New creates a new matrix with r rows and c columns
func New(r int, c int) Matrix {
	mat := make([][]int, r)
	for i := range mat {
		mat[i] = make([]int, c)
	}
	return Matrix{data: mat}
}

func isvalid(m1 Matrix, m2 Matrix) (valid bool) {
	m1Lenr, m1Lenc, m2Lenr, m2Lenc := len(m1.data), len(m1.data[0]), len(m2.data), len(m2.data[0])
	valid = (m1Lenr == m2Lenr && m1Lenc == m2Lenc)
	return
}

func iterate(m1 Matrix, operation func(int, int)) {
	for i, m1row := range m1.data {
		for j := range m1row {
			operation(i, j)
		}
	}
}

// Add adds two matrices if their dimensions are the same
// It returns a new matrix with the addition results,
// and the validity of the operation
func Add(m1 Matrix, m2 Matrix) (res Matrix, valid bool) {
	valid = isvalid(m1, m2)
	if valid {
		res = New(len(m1.data), len(m1.data[0]))

		iterate(m1, func(i int, j int) {
			res.data[i][j] = m1.data[i][j] + m2.data[i][j]
		})
	}

	return
}

// Subtract subtracts two matrices, m2 from m1,  if their dimensions are the same
// It returns a new matrix with the addition results,
// and the validity of the operation
func Subtract(m1 Matrix, m2 Matrix) (res Matrix, valid bool) {
	valid = isvalid(m1, m2)
	if valid {
		res = New(len(m1.data), len(m1.data[0]))

		iterate(m1, func(i int, j int) {
			res.data[i][j] = m1.data[i][j] - m2.data[i][j]
		})
	}

	return
}

// String returns a string representation of the matrix
func String(m Matrix) string {
	result := ""

	iterate(m, func(i int, j int) {
		delim := "\t"
		if j == 0 {
			delim = "\n"
		}

		result = result + fmt.Sprintf("%s%4d", delim, m.data[i][j])

	})

	result += "\n"

	return result
}

func (m *Matrix) Size() (int, int) {
	return len(m.data), len(m.data[0])
}

func (m *Matrix) Set(x int, y int, value int) (result bool) {
	rows, cols := m.Size()
	if x < rows && x >= 0 && y < cols && y >= 0 {
		m.data[x][y] = value
		result = true
	}
	return
}

func (m *Matrix) Get(x int, y int) (result int) {
	rows, cols := m.Size()
	if x < rows && x >= 0 && y < cols && y >= 0 {
		result = m.data[x][y]
	}
	return
}

func (m Matrix) String() string {
	return String(m)
}

func (m *Matrix) Add(other Matrix) {
	if isvalid(*m, other) {
		iterate(*m, func(i int, j int) {
			m.data[i][j] += other.data[i][j]
		})
	}
}

func (m *Matrix) Subtract(other Matrix) {
	if isvalid(*m, other) {
		iterate(*m, func(i int, j int) {
			m.data[i][j] -= other.data[i][j]
		})
	}
}
