package math

func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func MinUint(a, b uint) uint {
	if a > b {
		return b
	}
	return a
}
