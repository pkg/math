package math

func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func MinUint16(a, b uint16) uint16 {
	if a > b {
		return b
	}
	return a
}
