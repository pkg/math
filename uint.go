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
	if a > b {
		return b
	}
	return a
}
