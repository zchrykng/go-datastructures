package containers

import (
	"cmp"
	"slices"
)

type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}

func GetSortedValues[T cmp.Ordered](container Container[T]) []T {
	values := container.Values()
	if len(values) < 2 {
		return values
	}

	slices.Sort(values)
	return values
}

func GetSortedValuesFunc[T any](container Container[T], comparator func(a, b T) int) []T {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	slices.SortFunc(values, comparator)
	return values
}
