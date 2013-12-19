package math

// MaxUint16 returns the largest uint16 in the set provided.
// If no values are provided, MaxUint16 returns 0.
func MaxUint16(v ...uint16) uint16 {
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

// MinUint16 returns the smallest uint16 in the set provided.
// If no values are provided, MinUint16 returns 0.
func MinUint16(v ...uint16) uint16 {
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
