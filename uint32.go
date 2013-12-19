package math

// MaxUint32 returns the largest uint32 in the set provided.
// If no values are provided, MaxUint32 returns 0.
func MaxUint32(v ...uint32) uint32 {
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

// MinUint32 returns the smallest uint32 in the set provided.
// If no values are provided, MinUint32 returns 0.
func MinUint32(v ...uint32) uint32 {
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
