package math

// MaxUint16 returns the largest uint16 in the set provided.
// If no values are provided, MaxUint16 returns 0.
func MaxUint16(v ...uint16) uint16 {
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
		return MaxUint16(MaxUint16(v[:l]...), MaxUint16(v[l:]...))
	}
}

// MinUint16 returns the smallest uint16 in the set provided.
// If no values are provided, MinUint16 returns 0.
func MinUint16(v ...uint16) uint16 {
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
		return MinUint16(MinUint16(v[:l]...), MinUint16(v[l:]...))
	}
}
