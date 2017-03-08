package algorithms

import (
	"errors"
	"fmt"
)

type Matrix struct {
	rows int
	cols int
	data [][]int
}

func (m *Matrix) Print() {
	str := ""
	for i := 0; i < m.rows; i++ {
		str += fmt.Sprintf("%5d\n", m.data[i])
	}
	fmt.Println(str)
}

func NewMartix(r, c int) *Matrix {
	mat := &Matrix{rows: r, cols: c}
	data := make([][]int, r)
	for i := 0; i < r; i++ {
		data[i] = make([]int, c)
	}
	mat.data = data
	return mat
}

//方矩阵乘法
func SquareMatrixMultiply(a, b *Matrix) *Matrix {
	c := NewMartix(a.rows, a.rows)

	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.rows; j++ {
			for k := 0; k < a.rows; k++ {
				c.data[i][j] = c.data[i][j] + a.data[i][k]*b.data[k][j]
			}
		}
	}
	return c
}

//矩阵乘法
func MatrixMultiply(a, b *Matrix) (*Matrix, error) {
	if a.cols != b.rows {
		return nil, errors.New("a.cols != b.rows")
	}

	c := NewMartix(a.rows, b.cols)
	for i := 0; i < a.rows; i++ {
		for j := 0; j < b.cols; j++ {
			for k := 0; k < a.cols; k++ {
				c.data[i][j] = c.data[i][j] + a.data[i][k]*b.data[k][j]
			}
		}
	}
	return c, nil
}

//Strassen算法矩阵乘法(2*2)(错误)
func StrassenSquareMatrixMultiply(a, b *Matrix) *Matrix {
	c := NewMartix(a.rows, a.rows)

	x1 := (a.data[0][0] + a.data[1][1]) * (b.data[0][0] + b.data[1][1])
	x2 := (a.data[1][0] + a.data[1][1]) * b.data[0][0]
	x3 := a.data[0][0] * (b.data[0][1] - b.data[1][1])
	x4 := a.data[1][1] * (b.data[1][0] - b.data[0][0])
	x5 := (a.data[0][0] + a.data[0][1]) * b.data[1][1]
	x6 := (a.data[1][0] - a.data[0][0]) * (b.data[0][0] + b.data[0][1])
	x7 := (a.data[0][1] - a.data[1][1]) * (b.data[1][0] + b.data[1][1])

	c.data[0][0] = x1 + x4 - x5 + x7
	c.data[0][1] = x3 + x5
	c.data[1][0] = x2 + x4
	c.data[1][1] = x1 - x2 + x3 + x6
	return c
}
