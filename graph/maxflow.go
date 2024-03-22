package graph

import (
	"fmt"
)

type MaxFlowRunner struct {
	totalNodes    int
	capacityGraph [][]int
}

func NewMaxFlowRunner(totalNodes int) *MaxFlowRunner {
	capacityGraph := make([][]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		capacityGraph[i] = make([]int, totalNodes)
	}

	return &MaxFlowRunner{
		totalNodes:    totalNodes,
		capacityGraph: capacityGraph,
	}
}

func (m *MaxFlowRunner) findPath(src, target int) (bool, []int) {
	queue := make([]int, 0)
	queue = append(queue, src)

	visited := make([]bool, m.totalNodes)
	visited[src] = true
	parents := make([]int, m.totalNodes)
	for i := 0; i < m.totalNodes; i++ {
		parents[i] = -1
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == target {
			// fmt.Println("found the target", parents)
			path := make([]int, 0)
			curr := node
			for curr != -1 {
				path = append(path, curr)
				curr = parents[curr]
			}
			reverseArr(path)

			return true, path
		}

		for i := 0; i < m.totalNodes; i++ {
			if !visited[i] && m.capacityGraph[node][i] > 0 {
				visited[i] = true
				queue = append(queue, i)
				parents[i] = node
			}
		}
	}

	return false, nil
}

func (m *MaxFlowRunner) Run(src, target int) int {
	totalFlow := 0
	for {
		// m.printGraph()
		// find the augmented path in the residual capacity graph
		found, path := m.findPath(src, target)

		if !found {
			break
		}

		// on the path, find the minimum capacity
		minCap := m.capacityGraph[path[0]][path[1]]
		for i := 2; i < len(path); i++ {
			if minCap > m.capacityGraph[path[i-1]][path[i]] {
				minCap = m.capacityGraph[path[i-1]][path[i]]
			}
		}
		// fmt.Println("path", path, "min", minCap)
		// if minCap == 0 {
		// 	panic("unexpected")
		// }

		// then update the augmented graph capacity
		for i := 1; i < len(path); i++ {
			from, to := path[i-1], path[i]

			m.capacityGraph[from][to] -= minCap // the forward edge
			m.capacityGraph[to][from] += minCap // the backward edge
		}
		totalFlow += minCap
	}
	return totalFlow
}

func (m *MaxFlowRunner) AddEdge(src, dest, cap int) {
	m.capacityGraph[src][dest] += cap
}

func (m *MaxFlowRunner) printGraph() {
	// fmt.Println("graph======")
	for i := 0; i < m.totalNodes; i++ {
		fmt.Println(m.capacityGraph[i])
	}
}

func reverseArr[T any](arr []T) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
