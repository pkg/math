package math

// MaxInt16 returns the largest int16 in the set provided.
// If no values are provided, MaxInt16 returns 0.
func MaxInt16(v ...int16) int16 {
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

// MinInt16 returns the smallest int16 in the set provided.
// If no values are provided, MinInt16 returns 0.
func MinInt16(v ...int16) int16 {
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
