package math

// MaxInt returns the larger of two ints.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt returns the smaller of two ints.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxIntN returns the largest int in the set provided.
// If no values are provided, MaxInt returns 0.
func MaxIntN(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxInt(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxIntN(MaxIntN(v[:l]...), MaxIntN(v[l:]...))
	}
}

// MinIntN returns the smallest int in the set provided.
// If no values are provided, MinInt returns 0.
func MinIntN(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinInt(v[0], v[1])
	default:
		l := len(v) / 2
		return MinIntN(MinIntN(v[:l]...), MinIntN(v[l:]...))
	}
}
