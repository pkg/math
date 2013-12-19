package math

// MaxInt32 returns the largest int32 in the set provided.
// If no values are provided, MaxInt32 returns 0.
func MaxInt32(v ...int32) int32 {
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

// MinInt32 returns the smallest int32 in the set provided.
// If no values are provided, MinInt32 returns 0.
func MinInt32(v ...int32) int32 {
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
