package arraystack

import "github.com/zchrykng/go-datastructures/containers"

var _ containers.JSONSerializer = (*Stack[int])(nil)
var _ containers.JSONDeserializer = (*Stack[int])(nil)

func (stack *Stack[T]) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

func (stack *Stack[T]) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}

func (stack *Stack[T]) UnmarshalJSON(bytes []byte) error {
	return stack.FromJSON(bytes)
}

func (stack *Stack[T]) MarshalJSON() ([]byte, error) {
	return stack.ToJSON()
}
