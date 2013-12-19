package math

// MaxUint returns the largest uint in the set provided.
// If no values are provided, MaxUint returns 0.
func MaxUint(v ...uint) uint {
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
		return MaxUint(MaxUint(v[:l]...), MaxUint(v[l:]...))
	}
}

// MinUint returns the smallest uint in the set provided.
// If no values are provided, MinUint returns 0.
func MinUint(v ...uint) uint {
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
		return MinUint(MinUint(v[:l]...), MinUint(v[l:]...))
	}
}
