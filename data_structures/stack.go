package datastructures

type stack[T any] struct {
	data []T
}

func NewStack[T any]() *stack[T] {
	s := stack[T]{}
	return &s
}

func (s *stack[T]) Stack(element T) {
	s.data = append([]T{element}, s.data...)
}

func (s *stack[T]) Destack() (element T, ok bool) {
	var empty T
	if len(s.data) <= 0 {
		return empty, false
	}
	d := s.data[0]
	s.data = s.data[1:]
	return d, true
}

func (s *stack[T]) Read() (element T, ok bool) {
	var empty T
	if len(s.data) <= 0 {
		return empty, false
	}
	return s.data[0], true
}
