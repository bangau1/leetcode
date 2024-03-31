package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonotonicDequeu(t *testing.T) {
	mq := NewMonotonicDequeue[int](func(a, b int) bool {
		return a < b
	})

	nums := []int{1, 2, -1, 10, 2}
	expected := []int{-1, 2}

	for _, num := range nums {
		mq.PushBack(num)
	}

	assert.Equal(t, len(expected), mq.Len())
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], mq.Front())
		mq.RemoveFront()
	}
}
