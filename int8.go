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
	if a < b {
		return a
	}
	return b
}

// MaxInt8N returns the largest int8 in the set provided.
// If no values are provided, MaxInt8 returns 0.
func MaxInt8N(v ...int8) int8 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxInt8(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxInt8N(MaxInt8N(v[:l]...), MaxInt8N(v[l:]...))
	}
}

// MinInt8N returns the smallest int8 in the set provided.
// If no values are provided, MinInt8 returns 0.
func MinInt8N(v ...int8) int8 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinInt8(v[0], v[1])
	default:
		l := len(v) / 2
		return MinInt8N(MinInt8N(v[:l]...), MinInt8N(v[l:]...))
	}
}
