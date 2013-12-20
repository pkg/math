package math

// MaxUint8 returns the larger of two uint8s.
func MaxUint8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

// MinUint8 returns the smaller of two uint8s.
func MinUint8(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

// MaxUint8N returns the largest uint8 in the set provided.
// If no values are provided, MaxUint8 returns 0.
func MaxUint8N(v ...uint8) uint8 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxUint8(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxUint8N(MaxUint8N(v[:l]...), MaxUint8N(v[l:]...))
	}
}

// MinUint8N returns the smallest uint8 in the set provided.
// If no values are provided, MinUint8 returns 0.
func MinUint8N(v ...uint8) uint8 {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinUint8(v[0], v[1])
	default:
		l := len(v) / 2
		return MinUint8N(MinUint8N(v[:l]...), MinUint8N(v[l:]...))
	}
}
