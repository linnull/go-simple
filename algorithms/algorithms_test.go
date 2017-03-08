package algorithms

import (
	"fmt"
	"github.com/linnull/goutil"
	"math/rand"
	"testing"
	"time"
)

var data = goutil.RandIntArray(21, 50)

func TestMaxCrossingSubArray(t *testing.T) {

	l, r, c := MaxCrossingSubArray(data, 0, len(data)/2, len(data)-1)

	fmt.Println(data)
	fmt.Printf("left%d\n", l)
	fmt.Printf("right%d\n", r)
	fmt.Printf("crossing%d\n", c)
}

func TestMaxSubArray(t *testing.T) {

	data := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, -15, -4, 7}
	// data := []int{13, 3, 25, 20, 3, 16, 23, 18, 20, 7, 12, 5, 22, 15, 4, 7}
	// data := []int{-13, -3, -25, -20, -3, -16, -23, -18, -20, -7, -12, -5, -22, -15, -4, -7}
	l, h, s := MaxSubArray(data, 0, len(data)-1)

	fmt.Println(data)
	fmt.Printf("low%d\n", l)
	fmt.Printf("high%d\n", h)
	fmt.Printf("sum%d\n", s)
}

func TestSquareMatrixMultiply(t *testing.T) {
	a := NewMartix(4, 4)
	b := NewMartix(4, 4)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = r.Intn(100)
			b.data[i][j] = r.Intn(100)
		}
	}

	c := SquareMatrixMultiply(a, b)

	a.Print()
	b.Print()
	c.Print()
}

func BenchmarkSquareMatrixMultiply(B *testing.B) {
	a := NewMartix(4, 4)
	b := NewMartix(4, 4)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = r.Intn(100)
			b.data[i][j] = r.Intn(100)
		}
	}
	for i := 0; i < B.N; i++ {
		SquareMatrixMultiply(a, b)
	}

}

func TestMatrixMultiply(t *testing.T) {
	a := NewMartix(3, 15)
	b := NewMartix(15, 4)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = r.Intn(50)
		}
	}
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			b.data[i][j] = r.Intn(50)
		}
	}

	c, err := MatrixMultiply(a, b)
	if err != nil {
		t.Fatal(err)
	}

	a.Print()
	b.Print()
	c.Print()
}

func BenchmarkMatrixMultiply(B *testing.B) {
	a := NewMartix(3, 15)
	b := NewMartix(15, 4)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = r.Intn(50)
		}
	}
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			b.data[i][j] = r.Intn(50)
		}
	}
	for i := 0; i < B.N; i++ {
		MatrixMultiply(a, b)
	}

}

func TestStrassenSquareMatrixMultiply(t *testing.T) {
	a := NewMartix(2, 2)
	b := NewMartix(2, 2)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = r.Intn(50)
		}
	}
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			b.data[i][j] = r.Intn(50)
		}
	}

	c := StrassenSquareMatrixMultiply(a, b)

	a.Print()
	b.Print()
	c.Print()
}

func BenchmarkStrassenSquareMatrixMultiply(B *testing.B) {
	a := NewMartix(2, 2)
	b := NewMartix(2, 2)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = r.Intn(50)
		}
	}
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			b.data[i][j] = r.Intn(50)
		}
	}
	for i := 0; i < B.N; i++ {
		StrassenSquareMatrixMultiply(a, b)
	}

}
