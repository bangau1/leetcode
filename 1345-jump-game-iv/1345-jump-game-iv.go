func minJumps(arr []int) int {
    // this can be solved by using priority queue (djikstra)
    n := len(arr)
    pos := make(map[int][]int)
    
    for idx, num := range arr {
        pos[num] = append(pos[num], idx)
    }

    iterateFunc := func(i int) []int {
        curr := arr[i]
        var res []int
        if i - 1 >= 0 {
           res = append(res, i-1)
        }
        if i + 1 < n {
           res = append(res, i+1)
        }
        for _, idx := range pos[curr]{
            if idx < i-1 || idx > i+1{
                res = append(res, idx)
            }
        }
        return res
    }
    removeFunc := func(i int) {
        curr := arr[i]
        delete(pos, curr)
    }
   
    return djikstra(0, n, iterateFunc, removeFunc)
}

func djikstra(source int, n int, iterateFunc func(idx int)[]int, removeFunc func(idx int)) int{
    
    dist := make([]int, n)
    for i:=0;i<n;i++{
        dist[i] = math.MaxInt
    }
    dist[source] = 0
    visited := make([]bool, n)
    pq := make([]vertex, 0)
    pq = append(pq, vertex{source, 0})
    for len(pq) > 0 {
        u := pq[0]
        for len(pq) > 0 && visited[u.index] {
            pq = pq[1:]
            u = pq[0]
        }
        
        if visited[u.index] {
            continue
        }
        if u.index == n-1{
            return u.jump
        }

        visited[u.index] = true
        for _, next := range iterateFunc(u.index){
            if !visited[next] && dist[u.index] + 1 < dist[next]{
                dist[next] = dist[u.index] + 1
                pq = append(pq, vertex{next, dist[next]})
            }
            removeFunc(u.index)
        }
    }
    return dist[n-1]
}

func Less(a, b vertex) bool {
    return a.jump < b.jump
}

type vertex struct {
    index int
    jump int
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
