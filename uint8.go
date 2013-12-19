package math

// MaxUint8 returns the largest uint8 in the set provided.
// If no values are provided, MaxUint8 returns 0.
func MaxUint8(v ...uint8) uint8 {
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
		return MaxUint8(MaxUint8(v[:l]...), MaxUint8(v[l:]...))
	}
}

// MinUint8 returns the smallest uint8 in the set provided.
// If no values are provided, MinUint8 returns 0.
func MinUint8(v ...uint8) uint8 {
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
		return MinUint8(MinUint8(v[:l]...), MinUint8(v[l:]...))
	}
}
