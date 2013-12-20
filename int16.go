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
	if a < b {
		return a
	}
	return b
}

// MaxInt16N returns the largest int16 in the set provided.
// If no values are provided, MaxInt16 returns 0.
func MaxInt16N(v ...int16) int16 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxInt16(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxInt16N(MaxInt16N(v[:l]...), MaxInt16N(v[l:]...))
	}
}

// MinInt16N returns the smallest int16 in the set provided.
// If no values are provided, MinInt16 returns 0.
func MinInt16N(v ...int16) int16 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinInt16(v[0], v[1])
	default:
		l := len(v) / 2
		return MinInt16N(MinInt16N(v[:l]...), MinInt16N(v[l:]...))
	}
}
