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
	if a > b {
		return b
	}
	return a
}
