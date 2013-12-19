package math

// MaxUint8 returns the largest uint8 in the set provided.
// If no values are provided, MaxUint8 returns 0.
func MaxUint8(v ...uint8) uint8 {
	if len(v) == 0 {
		return 0
	}
	res := v[0]
	for _, i := range v[1:] {
		if i >= res {
			res = i
		}
	}
	return res
}

// MinUint8 returns the smallest uint8 in the set provided.
// If no values are provided, MinUint8 returns 0.
func MinUint8(v ...uint8) uint8 {
	if len(v) == 0 {
		return 0
	}
	res := v[0]
	for _, i := range v[1:] {
		if i < res {
			res = i
		}
	}
	return res
}
