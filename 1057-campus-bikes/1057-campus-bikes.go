type pair struct {
    worker, bike int
    dist int
}

type minHeap[T any] struct{
    data []T
    less func(a, b T) bool
} 

func NewMinHeap[T any](less func(a, b T)bool, initial ...T) minHeap[T] {
    res := minHeap[T]{
        less: less,
    }
    if len(initial) > 0 {
        res.data = make([]T, len(initial))
        copy(res.data, initial)
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
    m.data = m.data[0:l-1]
    return item
}

func (m *minHeap[T]) Push(a any) {
    m.data = append(m.data, a.(T))
}

func assignBikes(workers [][]int, bikes [][]int) []int {
    pqList := NewMinHeap[minHeap[pair]](LessPqFunc)

    for i, workerPos := range workers {
        pqWorker := NewMinHeap[pair](LessPairFunc)
        for ii, bikePos := range bikes {
            dist := abs(workerPos[0] - bikePos[0]) + abs(workerPos[1] - bikePos[1])
            candidate := pair{
                i,ii, dist,
            }
            pqWorker.data = append(pqWorker.data, candidate)
            
        }
        heap.Init(&pqWorker)
        pqList.data = append(pqList.data, pqWorker)
    }
    heap.Init(&pqList)
    res := make([]int, len(workers))
    bikeFlag := make([]bool, len(bikes))
    
    for pqList.Len() > 0 {
        pairData := pqList.data[0].data[0]
        for bikeFlag[pairData.bike] {
            heap.Pop(&(pqList.data[0]))
            heap.Fix(&pqList, 0)
            pairData = pqList.data[0].data[0]
        }
        res[pairData.worker] = pairData.bike
        bikeFlag[pairData.bike] = true
        heap.Pop(&pqList)
    }
    return res

}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func LessPairFunc(a, b pair) bool {
    if a.dist == b.dist {
        if a.worker == b.worker {
            return a.bike < b.bike
        }
        return a.worker < b.worker
    }
    return a.dist < b.dist
}

func LessPqFunc(a, b minHeap[pair]) bool {
    return LessPairFunc(a.data[0], b.data[0])
}

func normalPqSolution(workers [][]int, bikes [][]int) []int{
    pq := NewMinHeap[pair](LessPairFunc)
    
    for i, workerPos := range workers {
        for ii, bikePos := range bikes {
            dist := abs(workerPos[0] - bikePos[0]) + abs(workerPos[1] - bikePos[1])
            candidate := pair{
                i,ii, dist,
            }
            pq.data = append(pq.data, candidate)
            
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