package math

// MaxUint returns the largest uint in the set provided.
// If no values are provided, MaxUint returns 0.
func MaxUint(v ...uint) uint {
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

// MinUint returns the smallest uint in the set provided.
// If no values are provided, MinUint returns 0.
func MinUint(v ...uint) uint {
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
