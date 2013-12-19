package math

// MaxInt64 returns the largest int64 in the set provided.
// If no values are provided, MaxInt64 returns 0.
func MaxInt64(v ...int64) int64 {
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
		return MaxInt64(MaxInt64(v[:l]...), MaxInt64(v[l:]...))
	}
}

// MinInt64 returns the smallest int64 in the set provided.
// If no values are provided, MinInt64 returns 0.
func MinInt64(v ...int64) int64 {
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
		return MinInt64(MinInt64(v[:l]...), MinInt64(v[l:]...))
	}
}
