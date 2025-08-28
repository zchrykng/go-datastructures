package singlylinkedlist

import (
	"encoding/json"

	"github.com/zchrykng/go-datastructures/containers"
)

var _ containers.JSONSerializer = (*List[int])(nil)
var _ containers.JSONDeserializer = (*List[int])(nil)

func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

func (list *List[T]) FromJSON(data []byte) error {
	var elements []T
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}

func (list *List[T]) UnmarshalJSON(bytes []byte) error {
	return list.FromJSON(bytes)
}

func (list *List[T]) MarshalJSON() ([]byte, error) {
	return list.ToJSON()
}
