package math

// MaxUint32 returns the larger of two uint32s.
func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

// MinUint32 returns the smaller of two uint32s.
func MinUint32(a, b uint32) uint32 {
	if a > b {
		return b
	}
	return a
}
