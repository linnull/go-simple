package algorithms

import (
	"math"
)

//插入排序
func InsertionSort(arr []int) {
	var i, j int
	var key int
	for i = 1; i < len(arr); i++ {
		key = arr[i]
		for j = i; j > 0 && key < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = key
	}
}

//归并排序
func Mergesort(arr []int) {
	mergesort(arr, 0, len(arr)-1)
}

//归并排序递归函数
func mergesort(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergesort(arr, p, q)
		mergesort(arr, q+1, r)
		merge1(arr, p, q, r)
	}
}

//归并排序中合并两个数组的函数（哨兵版）
func merge1(arr []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q

	left := make([]int, n1+1)
	copy(left, arr[p:q+1])
	left[n1] = math.MaxInt64

	right := make([]int, n2+1)
	copy(right, arr[q+1:r+1])
	right[n2] = math.MaxInt64

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}

//冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}
	}
}

//堆排序（最大堆）
func HeapSort(arr []int) {
	a := make([]int, 1)
	a = append(a, arr...)
	buildHeap(a)
	for i := len(a) - 1; i > 1; i-- {
		a[i], a[1] = a[1], a[i]
		maxHeapify(a[:i], 1)
	}
	copy(arr, a[1:])
}

//堆化
func maxHeapify(arr []int, i int) {
	l := 2 * i   //左儿子
	r := 2*i + 1 //右儿子
	var largest int
	if l <= len(arr)-1 && arr[l] > arr[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= len(arr)-1 && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest)
	}
}

//建堆（最大堆）
func buildHeap(arr []int) {
	for i := (len(arr) - 1) / 2; i >= 1; i-- {
		maxHeapify(arr, i)
	}
}

//快速排序
func Quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid, i := arr[0], 1
	head, tail := 0, len(arr)-1
	for head < tail {
		if arr[i] > mid {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	arr[head] = mid
	Quicksort(arr[:head])
	Quicksort(arr[head+1:])
}
