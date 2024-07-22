package cache

import "encoding/json"

// type Unmarshaler interface {
// 	UnmarshalJSON([]byte) error
// }

type Serializer[T any] interface {
	Unmarshal(data []byte, v *T) error
	Marshal(v *T) ([]byte, error)
}

type JSONSerializer[T any] struct {
}

func (s *JSONSerializer[T]) Unmarshal(data []byte, v *T) error {
	return json.Unmarshal(data, v)
}
func (s *JSONSerializer[T]) Marshal(v *T) ([]byte, error) {
	return json.Marshal(v)
}

func NewJSONSerializer[T any]() Serializer[T] {
	s := JSONSerializer[T]{}
	return &s
}
