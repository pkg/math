package math

// MaxInt8 returns the larger of two int8s.
func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

// MinInt8 returns the smaller of two int8s.
func MinInt8(a, b int8) int8 {
	if a > b {
		return b
	}
	return a
}
