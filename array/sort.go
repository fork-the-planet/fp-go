package array

import (
	G "github.com/ibm/fp-go/array/generic"
	O "github.com/ibm/fp-go/ord"
)

// Sort implements a stable sort on the array given the provided ordering
func Sort[T any](ord O.Ord[T]) func(ma []T) []T {
	return G.Sort[[]T](ord)
}
