package matrix

import "math/rand"

func GenerateMatrix(r int, c int) ([][]int) {
	mat := make([][]int, r)
	for i:=0; i<r; i++ {
		mat[i] = make([]int, c)
	}
	return mat
}

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

func Add(m1 [][]int, m2 [][]int) ([][]int, bool){
	m1_len_r := len(m1)
	m1_len_c := len(m1[0])
	m2_len_r := len(m2)
	m2_len_c := len(m2[0])
	res := GenerateMatrix(m1_len_r, m1_len_c)
	var valid bool
	valid = true
	// fmt.Println(m1_len_r, m2_len_r, m1_len_c, m2_len_c)
	if m1_len_r != m2_len_r || m1_len_c != m2_len_c {
		valid = false
	}
	if valid {
		for i:=0; i<m1_len_r; i++ {
			for j:=0; j<m1_len_c; j++ {
				res[i][j] = m1[i][j] + m2[i][j]
			}
		}
	}
	return res, valid
}

func Subtract(m1 [][]int, m2 [][]int) ([][]int, bool){
	m1_len_r := len(m1)
	m1_len_c := len(m1[0])
	m2_len_r := len(m2)
	m2_len_c := len(m2[0])
	res := GenerateMatrix(m1_len_r, m1_len_c)
	var valid bool
	valid = true
	// fmt.Println(m1_len_r, m2_len_r, m1_len_c, m2_len_c)
	if m1_len_r != m2_len_r || m1_len_c != m2_len_c {
		valid = false
	}
	if valid {
		for i:=0; i<m1_len_r; i++ {
			for j:=0; j<m1_len_c; j++ {
				res[i][j] = m1[i][j] - m2[i][j]
			}
		}
	}
	return res, valid
}

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