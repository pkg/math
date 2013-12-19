package math

// MaxInt returns the larger of two ints.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt returns the smaller of two ints.
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
