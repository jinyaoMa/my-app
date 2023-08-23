package funcs

// Any check if a slice met a condition
func Any[T comparable](slice []T, condition func(e T) bool) bool {
	for _, e := range slice {
		if condition(e) {
			return true
		}
	}
	return false
}

// Contains check if a slice contains an element
func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
