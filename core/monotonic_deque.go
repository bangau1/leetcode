package core

type MonotonicDequeu[T any] struct {
	data     []T
	lessFunc func(a, b T) bool
}

func NewMonotonicDequeue[T any](lessFunc func(a, b T) bool) *MonotonicDequeu[T] {
	return &MonotonicDequeu[T]{
		data:     make([]T, 0),
		lessFunc: lessFunc,
	}
}

func (m *MonotonicDequeu[T]) PushBack(val T) {
	for m.Len() > 0 && m.lessFunc(val, m.Back()) {
		m.RemoveBack()
	}
	m.data = append(m.data, val)
}

func (m *MonotonicDequeu[T]) Front() T {
	return m.data[0]
}

func (m *MonotonicDequeu[T]) Back() T {
	return m.data[len(m.data)-1]
}

func (m *MonotonicDequeu[T]) RemoveFront() {
	m.data = m.data[1:]
}

func (m *MonotonicDequeu[T]) RemoveBack() {
	m.data = m.data[:len(m.data)-1]
}

func (m *MonotonicDequeu[T]) Len() int {
	return len(m.data)
}
