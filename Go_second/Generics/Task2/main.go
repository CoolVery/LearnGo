package main

func Filter[T any](arr []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, a := range arr {
		if predicate(a) {
			result = append(result, a)
		}
	}
	return result
}
