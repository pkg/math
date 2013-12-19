package math

// MaxInt16 returns the largest int16 in the set provided.
// If no values are provided, MaxInt16 returns 0.
func MaxInt16(v ...int16) int16 {
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
		return MaxInt16(MaxInt16(v[:l]...), MaxInt16(v[l:]...))
	}
}

// MinInt16 returns the smallest int16 in the set provided.
// If no values are provided, MinInt16 returns 0.
func MinInt16(v ...int16) int16 {
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
		return MinInt16(MinInt16(v[:l]...), MinInt16(v[l:]...))
	}
}
