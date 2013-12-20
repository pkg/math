package math

// MaxUint32 returns the larger of two uint32s.
func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

// MinUint32 returns the smaller of two uint32s.
func MinUint32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

// MaxUint32N returns the largest uint32 in the set provided.
// If no values are provided, MaxUint32 returns 0.
func MaxUint32N(v ...uint32) uint32 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxUint32(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxUint32N(MaxUint32N(v[:l]...), MaxUint32N(v[l:]...))
	}
}

// MinUint32N returns the smallest uint32 in the set provided.
// If no values are provided, MinUint32 returns 0.
func MinUint32N(v ...uint32) uint32 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinUint32(v[0], v[1])
	default:
		l := len(v) / 2
		return MinUint32N(MinUint32N(v[:l]...), MinUint32N(v[l:]...))
	}
}
