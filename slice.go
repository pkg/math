package math

import "reflect"

// MinSliceIndex returns the smallest element index in the slice.
// If slice is empty, MinSliceIndex returns -1
func MinSliceIndex(slice interface{}, less func(i, j int) bool) int {
	v := reflect.ValueOf(slice)
	if v.Len() == 0 {
		return -1
	}
	minIndex := 0
	for i := 1; i < v.Len(); i++ {
		if less(i, minIndex) {
			minIndex = i
		}
	}
	return minIndex
}

// MaxSliceIndex returns the largest element index in the slice.
// If slice is empty, MaxSliceIndex returns -1
func MaxSliceIndex(slice interface{}, less func(i, j int) bool) int {
	v := reflect.ValueOf(slice)
	if v.Len() == 0 {
		return -1
	}
	maxIndex := 0
	for i := 1; i < v.Len(); i++ {
		if less(maxIndex, i) {
			maxIndex = i
		}
	}
	return maxIndex
}
