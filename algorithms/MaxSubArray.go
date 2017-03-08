package algorithms

import (
	"math"
)

const (
	Infinity = math.MaxInt32
)

//最大子数组
func MaxSubArray(array []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, array[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := MaxSubArray(array, low, mid)
		rightLow, rightHigh, rightSum := MaxSubArray(array, mid+1, high)
		crossLow, corssHigh, crossSum := MaxCrossingSubArray(array, low, mid, high)

		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, corssHigh, crossSum
		}
	}
}

//过中点最大子数组
func MaxCrossingSubArray(array []int, low, mid, high int) (maxLeft, maxRight, maxCrossing int) {
	leftSum := math.MinInt32
	sum := 0
	for i := mid; i >= low; i-- {
		sum += array[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += array[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	maxCrossing = leftSum + rightSum
	return
}
