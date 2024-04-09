func networkDelayTime(times [][]int, n int, k int) int {
    // we can use djikstra here, to calculate the shortest distance from node k to all other nodes
    // dist = djikstra(k, graph)
    // answer = max(dist)
    adjList := make([][]Vertex, n)
    for _, time := range times { // --> this is O(E), E = number of edges
        s, d, cost := time[0]-1, time[1]-1, time[2]
        adjList[s] = append(adjList[s], Vertex{
            d, cost,
        })
    }

    dist := djikstra(k-1, adjList)
    var res = 0
    for i:=0;i<n;i++{ // --> this is O(V), where V = number of nodes
        if dist[i] == math.MaxInt {
            return -1
        }

        if res < dist[i] {
            res = dist[i]
        }
    }
    return res
}

func djikstra(source int, adjList [][]Vertex) []int {
    n := len(adjList)
    dist := make([]int, n)
    visited := make([]bool, n)
    for i:=0;i<n;i++{ // --> O(V)
        dist[i] = math.MaxInt
    }
    dist[source] = 0

    pq := NewMinHeap[Vertex](LessFunc)
    heap.Push(pq, Vertex{source, 0})

    for pq.Len() > 0 { // --> this is we try to iterate all of the graph, the edges dynamically generated from each iteration from the binaryHeap
        source := heap.Pop(pq).(Vertex) // -> this is O(log V)

        if visited[source.node]{
            continue
        }

        visited[source.node] = true
        
        // visit all the neighborhood
        for _, next := range adjList[source.node] { // for each node, we try to get all the edges to inside the heap. This is O(E)
            if !visited[next.node] && source.cost + next.cost < dist[next.node] {
                dist[next.node] = source.cost + next.cost
                heap.Push(pq, Vertex{ // -> O(Log V)
                    next.node,
                    dist[next.node],
                })
            } 
        }
    }
    // in summary it's gonna be O(V + E log V)

    return dist
}


type Vertex struct {
    node int
    cost int
}

func LessFunc(a, b Vertex) bool {
    return a.cost < b.cost
}

type MinHeap[T any] struct {
    data []T
    lessFunc func(a, b T)bool
}

func NewMinHeap[T any](lessFunc func(a, b T) bool, data ...T) *MinHeap[T] {
    res := MinHeap[T]{
        lessFunc: lessFunc,
        data: make([]T, len(data)),
    }

    if len(data) > 0 {
        copy(res.data, data)
        heap.Init(&res)
    }
    return &res
    
}

func (m MinHeap[T]) Len() int{
    return len(m.data)
}

func (m MinHeap[T]) Less(a, b int) bool {
    return m.lessFunc(m.data[a], m.data[b])
}

func (m MinHeap[T]) Swap(a, b int) {
    m.data[a], m.data[b] = m.data[b], m.data[a]
}

func (m *MinHeap[T]) Pop() any {
    l := m.Len()
    item := m.data[l-1]
    m.data = m.data[:l-1]
    return item
}

func (m *MinHeap[T]) Push(a any) {
    m.data =append(m.data, a.(T))
}