package algorithms

import (
	// "fmt"
	"testing"

	"github.com/linnull/goutil"
)

func TestInsertionSort(t *testing.T) {
	var a0 = goutil.RandIntArray(10, 1000000)

	// fmt.Println(a0)
	InsertionSort(a0)
	// fmt.Println(a0)
}

func BenchmarkInsertionSort(b *testing.B) {
	// var a0 = goutil.RandIntArray(10, 1000000)
	// var a1 = goutil.RandIntArray(100, 1000000)
	// var a2 = goutil.RandIntArray(1000, 1000000)
	// var a3 = goutil.RandIntArray(10000, 1000000)
	var a4 = goutil.RandIntArray(100000, 1000000)

	for i := 0; i < b.N; i++ {
		// InsertionSort(a0)
		// InsertionSort(a1)
		// InsertionSort(a2)
		// InsertionSort(a3)
		InsertionSort(a4)
	}
}

func TestMergesort(t *testing.T) {
	var a0 = goutil.RandIntArray(10, 1000000)

	// fmt.Println(a0)
	Mergesort(a0)
	// fmt.Println(a0)
}

func BenchmarkMergesort(b *testing.B) {
	// var a0 = goutil.RandIntArray(10, 1000000)
	// var a1 = goutil.RandIntArray(100, 1000000)
	// var a2 = goutil.RandIntArray(1000, 1000000)
	// var a3 = goutil.RandIntArray(10000, 1000000)
	var a4 = goutil.RandIntArray(100000, 1000000)

	for i := 0; i < b.N; i++ {
		// Mergesort(a0)
		// Mergesort(a1)
		// Mergesort(a2)
		// Mergesort(a3)
		Mergesort(a4)
	}
}

func TestBubbleSort(t *testing.T) {
	var a0 = goutil.RandIntArray(10, 1000000)

	// fmt.Println(a0)
	BubbleSort(a0)
	// fmt.Println(a0)
}

func BenchmarkBubbleSort(b *testing.B) {
	// var a0 = goutil.RandIntArray(10, 1000000)
	// var a1 = goutil.RandIntArray(100, 1000000)
	// var a2 = goutil.RandIntArray(1000, 1000000)
	// var a3 = goutil.RandIntArray(10000, 1000000)
	// var a4 = goutil.RandIntArray(100000, 1000000)

	for i := 0; i < b.N; i++ {
		// BubbleSort(a0)
		// BubbleSort(a1)
		// BubbleSort(a2)
		// BubbleSort(a3)
		// BubbleSort(a4)
	}
}

func TestHeapSort(t *testing.T) {
	var a0 = goutil.RandIntArray(10, 1000000)
	// var a4 = goutil.RandIntArray(100000, 1000000)

	// fmt.Println(a0)
	HeapSort(a0)
	// fmt.Println(a0)
}

func BenchmarkHeapSort(b *testing.B) {
	// var a0 = goutil.RandIntArray(10, 1000000)
	// var a1 = goutil.RandIntArray(100, 1000000)
	// var a2 = goutil.RandIntArray(1000, 1000000)
	// var a3 = goutil.RandIntArray(10000, 1000000)
	var a4 = goutil.RandIntArray(100000, 1000000)

	for i := 0; i < b.N; i++ {
		// BubbleSort(a0)
		// BubbleSort(a1)
		// BubbleSort(a2)
		// BubbleSort(a3)
		HeapSort(a4)
	}
}

func TestQuicksort(t *testing.T) {
	var a0 = goutil.RandIntArray(10, 1000000)

	// fmt.Println(a0)
	Quicksort(a0)
	// fmt.Println(a0)
}

func BenchmarkQuicksort(b *testing.B) {
	// var a0 = goutil.RandIntArray(10, 1000000)
	// var a1 = goutil.RandIntArray(100, 1000000)
	// var a2 = goutil.RandIntArray(1000, 1000000)
	// var a3 = goutil.RandIntArray(10000, 1000000)
	var a4 = goutil.RandIntArray(100000, 1000000)

	for i := 0; i < b.N; i++ {
		// Quicksort(a0)
		// Quicksort(a1)
		// Quicksort(a2)
		// Quicksort(a3)
		Quicksort(a4)
	}
}
