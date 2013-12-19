package math

// MaxUint64 returns the larger of two uint64s.
func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// MinUint64 returns the smaller of two uint64s.
func MinUint64(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}
