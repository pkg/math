package math

// MaxInt returns the largest int in the set provided.
// If no values are provided, MaxInt returns 0.
func MaxInt(v ...int) int {
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

// MinInt returns the smallest int in the set provided.
// If no values are provided, MinInt returns 0.
func MinInt(v ...int) int {
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
