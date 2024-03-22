package graph_test

import (
	"testing"

	. "github.com/bangau1/leetcode/graph"
	"github.com/stretchr/testify/assert"
)

func TestMaxFlowCalculation(t *testing.T) {
	runner := NewMaxFlowRunner(6)
	runner.AddEdge(0, 1, 16)
	runner.AddEdge(0, 2, 13)
	runner.AddEdge(1, 2, 10)
	runner.AddEdge(1, 3, 12)
	runner.AddEdge(2, 1, 4)
	runner.AddEdge(2, 4, 14)
	runner.AddEdge(3, 2, 9)
	runner.AddEdge(3, 5, 20)
	runner.AddEdge(4, 3, 7)
	runner.AddEdge(4, 5, 4)

	assert.Equal(t, 23, runner.Run(0, 5))
	assert.Equal(t, 0, runner.Run(0, 5)) // second time will produce 0 because we already cut all edges from source to sink

	runner = NewMaxFlowRunner(4)
	runner.AddEdge(0, 1, 3)
	runner.AddEdge(0, 2, 2)
	runner.AddEdge(1, 2, 5)
	runner.AddEdge(1, 3, 2)
	runner.AddEdge(2, 3, 3)

	assert.Equal(t, 5, runner.Run(0, 3))
	assert.Equal(t, 0, runner.Run(0, 3))
}
