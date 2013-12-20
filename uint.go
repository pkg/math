package math

// MaxUint returns the larger of two uints.
func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

// MinUint returns the smaller of two uints.
func MinUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// MaxUintN returns the largest uint in the set provided.
// If no values are provided, MaxUint returns 0.
func MaxUintN(v ...uint) uint {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MaxUint(v[0], v[1])
	default:
		l := len(v) / 2
		return MaxUintN(MaxUintN(v[:l]...), MaxUintN(v[l:]...))
	}
}

// MinUintN returns the smallest uint in the set provided.
// If no values are provided, MinUint returns 0.
func MinUintN(v ...uint) uint {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return MinUint(v[0], v[1])
	default:
		l := len(v) / 2
		return MinUintN(MinUintN(v[:l]...), MinUintN(v[l:]...))
	}
}
