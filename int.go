package math

// MaxInt returns the largest int in the set provided.
// If no values are provided, MaxInt returns 0.
func MaxInt(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		if v[0] > v[1] {
			return v[0]
		}
		return v[1]
	default:
		l := len(v) / 2
		return MaxInt(MaxInt(v[:l]...), MaxInt(v[l:]...))
	}
}

// MinInt returns the smallest int in the set provided.
// If no values are provided, MinInt returns 0.
func MinInt(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		if v[0] < v[1] {
			return v[0]
		}
		return v[1]
	default:
		l := len(v) / 2
		return MinInt(MinInt(v[:l]...), MinInt(v[l:]...))
	}
}
