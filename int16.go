package math

// MaxInt16 returns the larger of two int16s.
func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

// MinInt16 returns the smaller of two int16s.
func MinInt16(a, b int16) int16 {
	if a > b {
		return b
	}
	return a
}
