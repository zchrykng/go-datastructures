package singlylinkedlist

import (
	"fmt"
	"slices"
	"strings"

	"github.com/zchrykng/go-datastructures/lists"
	"github.com/zchrykng/go-datastructures/utils"
)

var _ lists.List[int] = (*List[int])(nil)

type List[T comparable] struct {
	first *element[T]
	last  *element[T]
	size  int
}

type element[T comparable] struct {
	value T
	next  *element[T]
}

func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List[T]) Add(values ...T) {
	for _, value := range values {
		newElement := &element[T]{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func (list *List[T]) Append(values ...T) {
	list.Add(values...)
}

func (list *List[T]) Prepend(values ...T) {
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &element[T]{value: values[v], next: list.first}
		list.first = newElement
		if list.size == 0 {
			list.last = newElement
		}
		list.size++
	}
}

func (list *List[T]) Get(index int) (value T, ok bool) {
	if !list.withinRange(index) {
		var t T
		return t, false
	}

	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}

	return element.value, true
}

func (list *List[T]) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeElement *element[T]
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
		beforeElement = element
	}

	if element == list.first {
		list.first = element.next
	}

	if element == list.last {
		list.last = beforeElement
	}

	if beforeElement != nil {
		beforeElement.next = element.next
	}

	element = nil

	list.size--
}

func (list *List[T]) Contains(values ...T) bool {
	if len(values) == 0 {
		return true
	}

	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List[T]) Values() []T {
	values := make([]T, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

func (list *List[T]) IndexOf(value T) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.Values() {
		if element == value {
			return index
		}
	}
	return -1
}

func (list *List[T]) Empty() bool {
	return list.size == 0
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if list.size < 2 {
		return
	}

	values := list.Values()
	slices.SortFunc(values, comparator)

	list.Clear()

	list.Add(values...)
}

func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var element1, element2 *element[T]
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

func (list *List[T]) Insert(index int, values ...T) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	if index == 0 {
		list.Prepend(values...)
		return
	}

	list.size += len(values)

	var beforeElement *element[T]
	foundElement := list.first
	for e := 0; e != index; e, foundElement = e+1, foundElement.next {
		beforeElement = foundElement
	}

	oldNextElement := beforeElement.next
	for _, value := range values {
		newElement := &element[T]{value: value}
		beforeElement.next = newElement
		beforeElement = newElement
	}
	beforeElement.next = oldNextElement
}

func (list *List[T]) Set(index int, value T) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}

	foundElement := list.first
	for e := 0; e != index; {
		e, foundElement = e+1, foundElement.next
	}
	foundElement.value = value
}

func (list *List[T]) String() string {
	str := "SinglyLinkedList\n"
	values := []string{}
	for element := list.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
