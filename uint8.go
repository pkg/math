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
	if a > b {
		return b
	}
	return a
}
