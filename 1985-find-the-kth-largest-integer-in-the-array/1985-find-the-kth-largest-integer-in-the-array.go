type MinHeap[T any] struct {
    data []T
    lessFunc func(a, b T) bool
}

func NewMinHeap[T any](lessFunc func(a, b T) bool, data ...T) MinHeap[T] {
    res := MinHeap[T]{
        data: make([]T, 0),
        lessFunc: lessFunc,
    }

    if len(data) > 0 {
        for _, datum := range data {
            res.data = append(res.data, datum)
        }
        heap.Init(&res)
    }

    return res
} 

func (m MinHeap[T]) Less(a, b int) bool {
    return m.lessFunc(m.data[a], m.data[b])
}

func (m MinHeap[T]) Swap(a, b int) {
    m.data[a], m.data[b] = m.data[b], m.data[a]
}

func (m MinHeap[T]) Len() int {
    return len(m.data)
}

func (m *MinHeap[T]) Push(a any) {
    m.data = append(m.data, a.(T))
}

func (m *MinHeap[T]) Pop() any {
    l := m.Len()
    item := m.data[l-1]
    m.data = m.data[:l-1]
    return item
}

func (m *MinHeap[T]) Peek() T {
    return m.data[0]
}

type PositiveBigInt string

func (p *PositiveBigInt) CompareTo(other PositiveBigInt) int {
    res := 0
    // defer func(){
    //     fmt.Println(*p, other, "result:", res)
    // }()
    pLen := len(*p)
    oLen := len(other)

    if pLen != oLen {
        res = pLen - oLen
        return res
    }

    // compare from the most significant bit
    for i:=0;i<pLen;i++{
        if (*p)[i] != (other)[i] {
            res = int((*p)[i]) - int((other)[i])
            return res
        }
    }

    return res
}


func kthLargestNumber(strs []string, k int) string {
    // to find kth largest number, we can use a priorityQueue or minHeap with size k (not maxHeap).
    // the idea is we compare the Peek() result and if the incoming data is greater than the Peek() (the min), then pop it 
    n := len(strs)
    nums := make([]PositiveBigInt, n)
    for i:=0;i<n;i++{
        nums[i] = PositiveBigInt(strs[i])
    }
    sort.Slice(nums, func (a, b int) bool {
        cmp := nums[a].CompareTo(nums[b])
        return cmp < 0
    })
    // fmt.Println("nums", nums)

    pq := NewMinHeap[PositiveBigInt](func (a, b PositiveBigInt) bool {
        cmp := a.CompareTo(b)
        return cmp < 0
    }, nums[0:k]...)

    for i:=k;i<n;i++{
        cmp := nums[i].CompareTo(pq.Peek())
        if cmp > 0 {
            pq.data[0] = nums[i]
            heap.Fix(&pq, 0)
        }
    }

    return string(pq.Peek())
}