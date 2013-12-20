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
	if a < b {
		return a
	}
	return b
}

// MaxInt32N returns the largest int32 in the set provided.
// If no values are provided, MaxInt32 returns 0.
func MaxInt32N(v ...int32) int32 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxInt32(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxInt32N(MaxInt32N(v[:l]...), MaxInt32N(v[l:]...))
	}
}

// MinInt32N returns the smallest int32 in the set provided.
// If no values are provided, MinInt32 returns 0.
func MinInt32N(v ...int32) int32 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinInt32(v[0], v[1])
	default:
		l := len(v) / 2
		return MinInt32N(MinInt32N(v[:l]...), MinInt32N(v[l:]...))
	}
}
