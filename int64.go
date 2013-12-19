package math

// MaxInt64 returns the largest int64 in the set provided.
// If no values are provided, MaxInt64 returns 0.
func MaxInt64(v ...int64) int64 {
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

// MinInt64 returns the smallest int64 in the set provided.
// If no values are provided, MinInt64 returns 0.
func MinInt64(v ...int64) int64 {
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
