package arraystack

import "github.com/zchrykng/go-datastructures/containers"

var _ containers.ReverseIteratorWithIndex[int] = (*Iterator[int])(nil)

type Iterator[T comparable] struct {
	stack *Stack[T]
	index int
}

func (stack *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{stack: stack, index: -1}
}

func (iterator *Iterator[T]) Next() bool {
	if iterator.index < iterator.stack.Size() {
		iterator.index++
	}
	return iterator.stack.withinRange(iterator.index)
}

func (iterator *Iterator[T]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.stack.withinRange(iterator.index)
}

func (iterator *Iterator[T]) Value() T {
	value, _ := iterator.stack.list.Get(iterator.stack.Size() - iterator.index - 1)
	return value
}

func (iterator *Iterator[T]) Index() int {
	return iterator.index
}

func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
}

func (iterator *Iterator[T]) End() {
	iterator.index = iterator.stack.Size()
}

func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator[T]) Last() bool {
	iterator.End()
	return iterator.Prev()
}

func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	for iterator.Next() {
		index, value := iterator.Index(), iterator.Value()
		if f(index, value) {
			return true
		}
	}
	return false
}

func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool {
	for iterator.Prev() {
		index, value := iterator.Index(), iterator.Value()
		if f(index, value) {
			return true
		}
	}

	return false
}
