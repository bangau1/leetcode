package graph_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bangau1/leetcode/graph"
	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	n := 10
	adjList := make([][]graph.Vertex, n)
	randObj := rand.New(rand.NewSource(time.Now().UnixMicro()))
	for i := 0; i < n; i++ {
		adjList[i] = make([]graph.Vertex, 0)
		for j := 0; j < n; j++ {
			if j != i {
				if randObj.Intn(100) >= 50 {
					continue
				}
				adjList[i] = append(adjList[i], graph.Vertex{
					Node: j,
					Cost: randObj.Intn(20),
				})
			}
		}
	}
	djikstraDist := graph.DjikstraAdjList(0, adjList)
	floydDist := graph.FloydWarshall(adjList)

	for i := 0; i < n; i++ {
		assert.Equal(t, djikstraDist[i], floydDist[0][i], "failed on i=%d", i)
	}

}
