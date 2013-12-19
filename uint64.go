package math

// MaxUint64 returns the largest uint64 in the set provided.
// If no values are provided, MaxUint64 returns 0.
func MaxUint64(v ...uint64) uint64 {
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

// MinUint64 returns the smallest uint64 in the set provided.
// If no values are provided, MinUint64 returns 0.
func MinUint64(v ...uint64) uint64 {
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
