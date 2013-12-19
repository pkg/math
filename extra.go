package math

// Max returns the larger of two ints.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two ints.
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
