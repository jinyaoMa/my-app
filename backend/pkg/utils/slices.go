package utils

func SliceMap[S ~[]E, E any, R any](slice S, m func(e E) R) (result []R) {
	for _, v := range slice {
		result = append(result, m(v))
	}
	return result
}

func SliceUnique[S ~[]E, E any](slice S, hash func(e E) string) (result []E) {
	result = make([]E, 0, len(slice))
	seen := make(map[string]bool)
	for _, val := range slice {
		id := hash(val)
		if _, ok := seen[id]; !ok {
			seen[id] = true
			result = append(result, val)
		}
	}
	return result
}

func SliceFilter[S ~[]E, E any](slice S, cond func(e E) bool) (result []E) {
	result = make([]E, 0, len(slice))
	for _, val := range slice {
		if cond(val) {
			result = append(result, val)
		}
	}
	return result
}

func SliceSum[S ~[]E, E int | int32 | int64 | uint | uint32 | uint64 | float32 | float64](slice S) (result E) {
	for _, v := range slice {
		result += v
	}
	return result
}
