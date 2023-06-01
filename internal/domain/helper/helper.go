package helper

import (
	"math"

	"golang.org/x/exp/constraints"
)

func FindMax[T constraints.Ordered](nums []T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}
	max := nums[0]
	for _, value := range nums {
		if max < value {
			max = value
		}
	}
	return max
}

func FindMin[T constraints.Ordered](nums []T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}
	min := nums[0]
	for _, value := range nums {
		if min > value {
			min = value
		}
	}
	return min
}

func CalculateEuclideanDistance[T int64 | int32 | float64 | float32](pointA, pointB []T) T {
	if len(pointA) != len(pointB) {
		var zero T
		return zero
	}
	var sum T
	for index := 0; index < len(pointA); index++ {
		sum += (pointA[index] - pointB[index]) * (pointA[index] - pointB[index])
	}
	return T(math.Sqrt(float64(sum)))
}
