package core_test

import (
	"container/heap"
	"testing"

	. "github.com/bangau1/leetcode/core"
	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	pq := NewMinHeap[int](func(a, b int) bool {
		return a < b
	}, []int{100, 3, 5}...)

	assert.Equal(t, 3, heap.Pop(pq))
	assert.Equal(t, 5, heap.Pop(pq))
	heap.Push(pq, -1)

	assert.Equal(t, -1, heap.Pop(pq))
	assert.Equal(t, 100, heap.Pop(pq))
	assert.Equal(t, 0, pq.Len())
}
