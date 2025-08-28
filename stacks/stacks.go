package stacks

import "github.com/zchrykng/go-datastructures/containers"

type Stack[T any] interface {
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
}
