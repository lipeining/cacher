package multicache

import jsoniter "github.com/json-iterator/go"

type Serializer[T any] interface {
	Unmarshal(data []byte, v *T) error
	Marshal(v *T) ([]byte, error)
}

type JSONSerializer[T any] struct {
}

func (s *JSONSerializer[T]) Unmarshal(data []byte, v *T) error {
	return jsoniter.Unmarshal(data, v)
}
func (s *JSONSerializer[T]) Marshal(v *T) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func NewJSONSerializer[T any]() Serializer[T] {
	s := JSONSerializer[T]{}
	return &s
}

type StringSerializer struct {
}

func (s *StringSerializer) Unmarshal(data []byte, v *string) error {
	*v = string(data)
	return nil
}

func (s *StringSerializer) Marshal(v *string) ([]byte, error) {
	return []byte(*v), nil
}

func NewStringSerializer() Serializer[string] {
	s := StringSerializer{}
	return &s
}
