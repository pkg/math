package math

// MaxInt8 returns the largest int8 in the set provided.
// If no values are provided, MaxInt8 returns 0.
func MaxInt8(v ...int8) int8 {
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
		return MaxInt8(MaxInt8(v[:l]...), MaxInt8(v[l:]...))
	}
}

// MinInt8 returns the smallest int8 in the set provided.
// If no values are provided, MinInt8 returns 0.
func MinInt8(v ...int8) int8 {
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
		return MinInt8(MinInt8(v[:l]...), MinInt8(v[l:]...))
	}
}
