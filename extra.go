package math

// Max returns the larger of two ints.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two ints.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxN returns the largest int in the set provided.
// If no values are provided, Max returns 0.
func MaxN(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return Max(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxN(MaxN(v[:l]...), MaxN(v[l:]...))
	}
}

// MinN returns the smallest int in the set provided.
// If no values are provided, Min returns 0.
func MinN(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return Min(v[0], v[1])
	default:
		l := len(v) / 2
		return MinN(MinN(v[:l]...), MinN(v[l:]...))
	}
}
