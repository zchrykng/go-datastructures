package arraystack

import (
	"fmt"
	"strings"

	"github.com/zchrykng/go-datastructures/lists/arraylist"
	"github.com/zchrykng/go-datastructures/stacks"
)

var _ stacks.Stack[int] = (*Stack[int])(nil)

type Stack[T comparable] struct {
	list *arraylist.List[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{list: arraylist.New[T]()}
}

func (stack *Stack[T]) Push(value T) {
	stack.list.Add(value)
}

func (stack *Stack[T]) Pop() (value T, ok bool) {
	value, ok = stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return
}

func (stack *Stack[T]) Peek() (value T, ok bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

func (stack *Stack[T]) Empty() bool {
	return stack.list.Empty()
}

func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}

func (stack *Stack[T]) Clear() {
	stack.list.Clear()
}

func (stack *Stack[T]) Values() []T {
	size := stack.list.Size()
	elements := make([]T, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1)
	}
	return elements
}

func (stack *Stack[T]) String() string {
	str := "ArrayStack\n"
	values := []string{}
	for _, value := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (stack *Stack[T]) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
