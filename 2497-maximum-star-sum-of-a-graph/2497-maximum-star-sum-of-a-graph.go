func maxStarSum(vals []int, edges [][]int, k int) int {
    if len(vals) == 1 {
        return vals[0]
    }
    if k == 0 {
        // only see the max node val
        res := vals[0]
        for _, val := range vals {
            if res < val {
                res = val
            }
        }
        return res
    }

    // a star graph that consists of n nodes:
    // - has a central node
    // - the neighbor has exactly 1 edge connecting to the central node
    // - has n-1 edge
    // - the leaf/neighbor nodes has 1 edge, while the star has n-1 edge

    // Given an integer k, return the maximum star sum of a star graph containing at most k edges.
    //
    // - manipulate the edges, so that the [a,b] -> a < b, otherwise swap it
    // - then sort the edges based on a
    // - then a becomes a central node
    // 
    // example:
    // for i:=0;i<len(edges);i++{
    //     // swap it
    //     if edges[i][0] > edges[i][1] {
    //         edges[i][0], edges[i][1] = edges[i][1], edges[i][0]
    //     }
    // }

    // sort.Slice(edges, func(a, b int) bool {
    //     return edges[a][0] < edges[b][0]
    // })

    // consider this sorted edges based on first node
    // - [0, 1], [0, 2], [0, 3], [1,5]
    // - there are 2 scenario:
    // 1. 0 is a central node, where the neighbor is 1,2,3. See that we can easily know the neighbor just by processing it from left to right where the firstnode=0
    // 2. 1 is a central node, the neighbor is 0, 5. See that left to right alone is not suffice. We should have mechanism to remember the neighbor in the left as well.
    // now we need to calculate the sum at most k edges, how we do that? where k <= 10^5-1.
    // 
    // another idea:
    // - should we have a minHeap for each node with size-k
    // - and we only put inside of it when the val is >= 0 (no negative value being put, except the central node).
    adjList := make([][]int, len(vals))
    for _, edge := range edges{
        a, b := edge[0], edge[1]
        adjList[a] = append(adjList[a], b)
        adjList[b] = append(adjList[b], a)
    }

    pqs := make([]*minHeap[int], len(vals))
    sum := make([]int, len(vals))
    for i:=0;i<len(vals);i++{
        sum[i] = vals[i] // the initial sum is the central node value
        pqs[i] = NewMinHeap[int](LessInt)
    }
    for i:=0;i<len(vals);i++{
        for _, adj := range adjList[i]{
            if vals[adj] >= 0 {
                ok, throw := addInt(pqs[i], vals[adj], k) // we can only add at most k neighbor value in the heap
                if ok{
                    sum[i] = sum[i] + vals[adj] - throw
                }
            }
        }
    }
    res := sum[0]
    // idx := -1
    for _, s := range sum {
        if res < s {
            res = s
            // idx = i
        }
    }
    // for i:=0;i<len(vals);i++{
    //     if pqs[i].Len() > k+1 {
    //         fmt.Println(pqs[i].data)
    //         panic("error")
    //     }
    // }
    // fmt.Println("max on central node", idx)
    // fmt.Println("heap on idx", pqs[idx].data)
    // fmt.Println("neighbor on idx", adjList[idx])

    return res

}

// we want to store only the max. 
func addInt(pq *minHeap[int], val, k int) (bool, int){
    if pq.Len() + 1 <= k {
        heap.Push(pq, val)
        return true, 0
    }else if pq.Peek() < val {
        pop := heap.Pop(pq).(int)
        heap.Push(pq, val)
        return true, pop
    }
    return false, -1
}

func LessInt(a, b int) bool {
    return a < b
}

type minHeap[T any] struct{
    data []T
    lessFunc func(a, b T) bool
}

func NewMinHeap[T any](lessFunc func(a, b T) bool, data ...T) *minHeap[T] {
    res := minHeap[T]{
        lessFunc: lessFunc,
    }
    if len(data) > 0 {
        res.data = make([]T, len(data))
        copy(res.data, data)
        heap.Init(&res)
    }else{
        res.data = make([]T, 0)
    }
    return &res
}

func (m minHeap[T]) Swap(a, b int) {
    m.data[a], m.data[b] = m.data[b], m.data[a]
}

func (m minHeap[T]) Less(a, b int) bool {
    return m.lessFunc(m.data[a], m.data[b])
}

func (m minHeap[T]) Len() int{
    return len(m.data)
}

func (m *minHeap[T]) Pop() any{
    l := m.Len()
    item := m.data[l-1]
    m.data = m.data[:l-1]
    return item
}

func (m *minHeap[T]) Peek() T {
    return m.data[0]
}

func (m *minHeap[T]) Push(a any) {
    m.data = append(m.data, a.(T))
}