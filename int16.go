package math

func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func MinInt16(a, b int16) int16 {
	if a > b {
		return b
	}
	return a
}
