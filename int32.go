package math

// MaxInt32 returns the largest int32 in the set provided.
// If no values are provided, MaxInt32 returns 0.
func MaxInt32(v ...int32) int32 {
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
		return MaxInt32(MaxInt32(v[:l]...), MaxInt32(v[l:]...))
	}
}

// MinInt32 returns the smallest int32 in the set provided.
// If no values are provided, MinInt32 returns 0.
func MinInt32(v ...int32) int32 {
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
		return MinInt32(MinInt32(v[:l]...), MinInt32(v[l:]...))
	}
}
