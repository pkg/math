package math

// MaxInt8 returns the largest int8 in the set provided.
// If no values are provided, MaxInt8 returns 0.
func MaxInt8(v ...int8) int8 {
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

// MinInt8 returns the smallest int8 in the set provided.
// If no values are provided, MinInt8 returns 0.
func MinInt8(v ...int8) int8 {
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
