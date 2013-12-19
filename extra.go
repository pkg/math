package math

// Max returns the largest int in the set provided.
// If no values are provided, Max returns 0.
func Max(v ...int) int {
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

// Min returns the smallest int in the set provided.
// If no values are provided, Min returns 0.
func Min(v ...int) int {
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
