type vertex struct{
    r,c1,c2 int
    total int
    prev *vertex
}

func greater(a, b vertex) bool {
    if a.r == b.r {
        // prioritize the bigger total r
        return a.total > b.total
    }
    return a.r < b.r
}

func cherryPickup(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    pq := NewMinHeap[vertex](greater, vertex{
        r: 0,
        c1: 0,
        c2: n-1,
        total: grid[0][0] + grid[0][n-1],
    })
    visited := make([][][]bool, m)
    for i:=0;i<m;i++{
        visited[i] = make([][]bool, n)
        for ii:=0;ii<n;ii++{
            visited[i][ii] = make([]bool, n)
        }
    }
    var last vertex
    var maxRow = 0
    for pq.Len() > 0 {
        node := heap.Pop(&pq).(vertex)
        if visited[node.r][node.c1][node.c2] {
            continue
        }
        if abs(maxRow-node.r) >= 2 {
            continue
        }
        maxRow = max(maxRow, node.r)
        if node.r == m-1 {
            last = node
            break
        }

        visited[node.r][node.c1][node.c2] = true
        for nc1 := node.c1-1;nc1<=node.c1+1;nc1++{
            for nc2:= node.c2-1;nc2<=node.c2+1;nc2++{
                
                if nc1 >= 0 && nc2 >= 0 && nc1 < n && nc2 < n && !visited[node.r+1][nc1][nc2]{
                    // visited[node.r+1][nc1][nc2] = true
                    total := node.total + grid[node.r+1][nc1]
                    if nc1 != nc2 {
                        total += grid[node.r+1][nc2]
                    }
                    heap.Push(&pq, vertex{node.r+1, nc1, nc2, total, &node})
                }
            }
        }
    }
    total := last.total
    // for true {
    //     fmt.Println(last.r, last.c1,last.c2, last.total)
    //     if last.prev == nil {
    //         break
    //     }
    //     last = *(last.prev)
        
    // }

    return total
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

// ==============
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
