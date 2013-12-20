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
	if a < b {
		return a
	}
	return b
}

// MaxInt64N returns the largest int64 in the set provided.
// If no values are provided, MaxInt64 returns 0.
func MaxInt64N(v ...int64) int64 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxInt64(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxInt64N(MaxInt64N(v[:l]...), MaxInt64N(v[l:]...))
	}
}

// MinInt64N returns the smallest int64 in the set provided.
// If no values are provided, MinInt64 returns 0.
func MinInt64N(v ...int64) int64 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinInt64(v[0], v[1])
	default:
		l := len(v) / 2
		return MinInt64N(MinInt64N(v[:l]...), MinInt64N(v[l:]...))
	}
}
