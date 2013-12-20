package math

// MaxUint64 returns the larger of two uint64s.
func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// MinUint64 returns the smaller of two uint64s.
func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// MaxUint64N returns the largest uint64 in the set provided.
// If no values are provided, MaxUint64 returns 0.
func MaxUint64N(v ...uint64) uint64 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxUint64(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxUint64N(MaxUint64N(v[:l]...), MaxUint64N(v[l:]...))
	}
}

// MinUint64N returns the smallest uint64 in the set provided.
// If no values are provided, MinUint64 returns 0.
func MinUint64N(v ...uint64) uint64 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinUint64(v[0], v[1])
	default:
		l := len(v) / 2
		return MinUint64N(MinUint64N(v[:l]...), MinUint64N(v[l:]...))
	}
}
