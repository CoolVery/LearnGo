package main

type NumConstraint interface {
	~int | ~float64 | ~float32
}

func Sum[T NumConstraint](nums[] T) T {
	var result T
	for _, n := range nums {
		result += n
	}
	return result
}