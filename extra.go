package math

// Max returns the largest int in the set provided.
// If no values are provided, Max returns 0.
func Max(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		if v[0] > v[1] {
			return v[0]
		}
		return v[1]
	default:
		l := len(v) / 2
		return Max(Max(v[:l]...), Max(v[l:]...))
	}
}

// Min returns the smallest int in the set provided.
// If no values are provided, Min returns 0.
func Min(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		if v[0] < v[1] {
			return v[0]
		}
		return v[1]
	default:
		l := len(v) / 2
		return Min(Min(v[:l]...), Min(v[l:]...))
	}
}
