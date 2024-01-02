type minHeap []vertex

func (m minHeap) Less(a, b int) bool {
    return m[a].dist < m[b].dist
}

func (m minHeap) Swap(a, b int) {
    m[a], m[b] = m[b], m[a]
}

func (m minHeap) Len() int {
    return len(m)
}

func (m *minHeap) Pop() any {
    l := m.Len()
    item := (*m)[l-1]
    *m = (*m)[0:l-1]
    return item
}

func (m *minHeap) Push(a any){
    *m = append(*m, a.(vertex))
}

func minimumDistance(n int, edges [][]int, s int, marked []int) int {
    pq := make(minHeap, 0)

    dist := make([]int, n)
    for i:=0;i<n;i++{
        dist[i] = math.MaxInt
    }

    adjMatrix := make([][]int, n)
    for i:=0;i<n;i++{
        adjMatrix[i] = make([]int, n)
        for j:=0;j<n;j++{
            adjMatrix[i][j] = math.MaxInt
        }
        adjMatrix[i][i] = 0
    }
    for _, edge := range edges {
        i, j, w := edge[0], edge[1], edge[2]
        adjMatrix[i][j] = min(adjMatrix[i][j], w)
    }
    visited := make([]bool, n)
    isMarkNode := make([]bool, n)
    for _, mark := range marked {
        isMarkNode[mark] = true
    }
    
    dist[s] = 0
    heap.Push(&pq, vertex{s, 0}) 
    for pq.Len() > 0 {
        u := heap.Pop(&pq).(vertex)
        if visited[u.node]{
            continue
        }

        visited[u.node] = true
        if isMarkNode[u.node] {
            return u.dist
        }

        for i:=0;i<n;i++{
            if u.node == i {
                continue
            }
            if adjMatrix[u.node][i] != math.MaxInt && !visited[i] && dist[i] > u.dist + adjMatrix[u.node][i] {
                dist[i] = u.dist + adjMatrix[u.node][i]
                heap.Push(&pq, vertex{i, dist[i]})
            }
        }
    }

    return -1
}

type vertex struct {
    node, dist int
}