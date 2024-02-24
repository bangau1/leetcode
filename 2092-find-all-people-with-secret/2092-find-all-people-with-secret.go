type edge struct {
    node int
    time int
}
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    // another approach to use:
    // djikstra with meetingTime as a weight
    // since time is chronological, then we must process the meeting that take places earlier
    // if somehow we can reach all of them earlier than the next meeting available, it means we don't need to go further
    adjList := make([][]edge, n)

    // fmt.Println("before")
    // printMatrix(matrix)
    adjList[0] = append(adjList[0], edge{firstPerson, 0})
    adjList[firstPerson] = append(adjList[firstPerson], edge{0, 0})
    
    for _, m := range meetings {
        a, b, time := m[0], m[1], m[2]
       
        adjList[a] = append(adjList[a], edge{b, time})
        adjList[b] = append(adjList[b], edge{a, time})
    }
    // fmt.Println("after")
    // printMatrix(matrix)
    var res []int
    dist := djikstra(0, adjList)
    // fmt.Println("dist", dist)
    for i:=0;i<n;i++{
        if dist[i] != math.MaxInt{
            res = append(res, i)
        }
    }
    return res
}

func printMatrix(matrix [][]int) {
    fmt.Println("===============")
    n := len(matrix)
    for i:=0;i<n;i++{
        fmt.Println(matrix[i])
    }
}


func djikstra(src int, adjList [][]edge) []int {
	n := len(adjList)
	dist := make([]int, n)
	visited := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
        visited[i] = make(map[int]bool)
	}
	dist[src] = 0
	pq := NewMinHeap[vertex](func(a, b vertex) bool {
		return a.time < b.time
	}, vertex{
		src, 0,
	})
    lastTime := 0
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(vertex)
        if u.node == 88 {
            fmt.Println("88")
        }
		if visited[u.node][u.time] {
			continue
		}
        if lastTime > u.time {
            panic("unexpected")
        }
        lastTime = max(u.time, lastTime)

		visited[u.node][u.time] = true
        for _, edge := range adjList[u.node] {
            next, time := edge.node, edge.time
            if u.time <= time {
                dist[next] = min(dist[next], time)
                heap.Push(&pq, vertex{next, dist[next]})
            } 
        }
	}

	return dist

}

type vertex struct {
    node int
    time int
}

type minHeap[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewMinHeap[T any](less func(a, b T) bool, data ...T) minHeap[T] {
	res := minHeap[T]{
		less: less,
	}

	if len(data) > 0 {
		res.data = make([]T, len(data))
		copy(res.data, data)
		heap.Init(&res)
	}

	return res
}

func (m minHeap[T]) Less(a, b int) bool {
	return m.less(m.data[a], m.data[b])
}

func (m minHeap[T]) Swap(a, b int) {
	m.data[a], m.data[b] = m.data[b], m.data[a]
}

func (m minHeap[T]) Len() int {
	return len(m.data)
}

func (m *minHeap[T]) Pop() any {
	l := m.Len()
	item := m.data[l-1]
	m.data = m.data[0 : l-1]
	return item
}

func (m *minHeap[T]) Push(a any) {
	m.data = append(m.data, a.(T))
}