package math

// MaxInt64 returns the larger of two int64s.
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MinInt64 returns the smaller of two int64s.
func MinInt64(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
