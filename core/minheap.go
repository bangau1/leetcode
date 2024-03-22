package core

import "container/heap"

type MinHeap[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewMinHeap[T any](less func(a, b T) bool, data ...T) *MinHeap[T] {
	res := MinHeap[T]{
		less: less,
	}

	if len(data) > 0 {
		res.data = make([]T, len(data))
		copy(res.data, data)
		heap.Init(&res)
	}

	return &res
}

func (m MinHeap[T]) Less(a, b int) bool {
	return m.less(m.data[a], m.data[b])
}

func (m MinHeap[T]) Swap(a, b int) {
	m.data[a], m.data[b] = m.data[b], m.data[a]
}

func (m MinHeap[T]) Len() int {
	return len(m.data)
}

func (m *MinHeap[T]) Pop() any {
	l := m.Len()
	item := m.data[l-1]
	m.data = m.data[0 : l-1]
	return item
}

func (m *MinHeap[T]) Push(a any) {
	m.data = append(m.data, a.(T))
}
