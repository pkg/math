package math

// MaxInt32 returns the larger of two int32s.
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// MinInt32 returns the smaller of two int32s.
func MinInt32(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}
