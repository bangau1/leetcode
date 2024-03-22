package graph

import (
	"container/heap"
	"math"

	"github.com/bangau1/leetcode/core"
)

type Vertex struct {
	Node, Cost int
}

// DjikstraAdjList with time complexity to : O(E Log V)
func DjikstraAdjList(src int, adjList [][]Vertex) []int {
	n := len(adjList)
	dist := make([]int, n)
	core.ArrayFill(dist, math.MaxInt)

	visited := make([]bool, n)

	pq := core.NewMinHeap[Vertex](func(a, b Vertex) bool {
		return a.Cost < b.Cost
	}, Vertex{src, 0}) // add source node
	dist[src] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Vertex)

		if visited[curr.Node] {
			continue
		}

		visited[curr.Node] = true

		for _, next := range adjList[curr.Node] {
			if !visited[next.Node] && next.Cost+curr.Cost < dist[next.Node] {
				dist[next.Node] = curr.Cost + next.Cost
				heap.Push(pq, Vertex{
					Node: next.Node,
					Cost: dist[next.Node],
				})
			}
		}
	}

	return dist
}

// All source shortest path
func FloydWarshall(adjList [][]Vertex) [][]int {
	n := len(adjList)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for ii := 0; ii < n; ii++ {
			dist[i][ii] = math.MaxInt
		}
		dist[i][i] = 0
	}

	for src, edges := range adjList {
		for _, edge := range edges {
			dist[src][edge.Node] = min(dist[src][edge.Node], edge.Cost)
		}
	}

	//loop naively via middle k idx
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
				if dist[i][k] != math.MaxInt && dist[k][j] != math.MaxInt && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}

			}
		}
	}
	return dist
}
