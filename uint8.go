package math

func MaxUint8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func MinUint8(a, b uint8) uint8 {
	if a > b {
		return b
	}
	return a
}
