package math

func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func MinUint32(a, b uint32) uint32 {
	if a > b {
		return b
	}
	return a
}
