package math

// MaxUint32 returns the largest uint32 in the set provided.
// If no values are provided, MaxUint32 returns 0.
func MaxUint32(v ...uint32) uint32 {
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
		return MaxUint32(MaxUint32(v[:l]...), MaxUint32(v[l:]...))
	}
}

// MinUint32 returns the smallest uint32 in the set provided.
// If no values are provided, MinUint32 returns 0.
func MinUint32(v ...uint32) uint32 {
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
		return MinUint32(MinUint32(v[:l]...), MinUint32(v[l:]...))
	}
}
