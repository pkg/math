package math

// MaxUint64 returns the largest uint64 in the set provided.
// If no values are provided, MaxUint64 returns 0.
func MaxUint64(v ...uint64) uint64 {
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
		return MaxUint64(MaxUint64(v[:l]...), MaxUint64(v[l:]...))
	}
}

// MinUint64 returns the smallest uint64 in the set provided.
// If no values are provided, MinUint64 returns 0.
func MinUint64(v ...uint64) uint64 {
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
		return MinUint64(MinUint64(v[:l]...), MinUint64(v[l:]...))
	}
}
