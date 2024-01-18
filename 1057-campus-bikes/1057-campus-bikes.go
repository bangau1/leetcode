type pair struct {
    worker, bike int
    dist int
}

type minHeap []pair

func (m minHeap) Less(a, b int) bool {
    if m[a].dist == m[b].dist {
        if m[a].worker == m[b].worker {
            return m[a].bike < m[b].bike
        }
        return m[a].worker < m[b].worker
    }
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

func (m *minHeap) Push(a any) {
    *m = append(*m, a.(pair))
}

func assignBikes(workers [][]int, bikes [][]int) []int {
    pq := make(minHeap, 0)
    
    for i, workerPos := range workers {
        for ii, bikePos := range bikes {
            dist := abs(workerPos[0] - bikePos[0]) + abs(workerPos[1] - bikePos[1])
            pq = append(pq, pair{
                i,ii, dist,
            })
        }
    }
    heap.Init(&pq)

    workerFlag := make([]bool, len(workers))
    bikeFlag := make([]bool, len(bikes))

    res := make([]int, len(workers))
    count := 0

    for count < len(workers) {
        data := heap.Pop(&pq).(pair)

        if !workerFlag[data.worker] && !bikeFlag[data.bike] {
            res[data.worker] = data.bike
            workerFlag[data.worker] = true
            bikeFlag[data.bike] = true
            count++
        }
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}