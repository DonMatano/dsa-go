package datastructures

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	q := Queue[T]{}
	return &q
}
func (q *Queue[T]) Enqueue(element T) {
	q.data = append(q.data, element)
}

func (q *Queue[T]) Dequeue() (element T, ok bool) {
	var empty T
	if len(q.data) <= 0 {
		return empty, false
	}
	d := q.data[0]
	q.data = q.data[1:]
	return d, true
}

func (q *Queue[T]) Read() (element T, ok bool) {
	var empty T
	if len(q.data) <= 0 {
		return empty, false
	}
	return q.data[0], true
}
