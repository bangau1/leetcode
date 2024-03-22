package graph_test

import (
	"testing"

	. "github.com/bangau1/leetcode/graph"
	"github.com/stretchr/testify/assert"
)

func TestDisjointSet(t *testing.T) {
	d := NewDisjointSet(3)
	assert.Equal(t, 3, d.GetTotalUnions())
	assert.Equal(t, 1, d.GetUnionSize(0))

	ok := d.Union(0, 1)
	assert.True(t, ok)
	assert.Equal(t, d.Find(0), d.Find(1))
	assert.NotEqual(t, d.Find(0), d.Find(2))
	assert.Equal(t, 2, d.GetTotalUnions())

	ok = d.Union(0, 1)
	assert.False(t, ok)

	ok = d.Union(0, 2)
	assert.True(t, ok)
	assert.Equal(t, 1, d.GetTotalUnions())
	assert.Equal(t, d.Find(0), d.Find(2))
}
