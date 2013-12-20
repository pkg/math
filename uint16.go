package math

// MaxUint16 returns the larger of two uint16s.
func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

// MinUint16 returns the smaller of two uint16s.
func MinUint16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

// MaxUint16N returns the largest uint16 in the set provided.
// If no values are provided, MaxUint16 returns 0.
func MaxUint16N(v ...uint16) uint16 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxUint16(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxUint16N(MaxUint16N(v[:l]...), MaxUint16N(v[l:]...))
	}
}

// MinUint16N returns the smallest uint16 in the set provided.
// If no values are provided, MinUint16 returns 0.
func MinUint16N(v ...uint16) uint16 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinUint16(v[0], v[1])
	default:
		l := len(v) / 2
		return MinUint16N(MinUint16N(v[:l]...), MinUint16N(v[l:]...))
	}
}
