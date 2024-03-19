func leastInterval(tasks []byte, n int) int {
    if n == 0 {
        return len(tasks)
    }
    // intuition:                              
    //
    // - keep a minHeap, prioritize by count
    // - also keep a queue with size = n + 1
    // - pop from minHeap,
    //    -  reduce the pop's count, increment the interval
    //    -  if pop's count == 0, do nothing
    //    - if pop's count > 0, then put it into the queue
    //      - if queue full, then remove queue's front
    //          - then add the minHeap's pop into the queue's back
    //          - the recently evicted item
    //             - if it's valid character, put back into the minHeap
    //             - otherwise no-op
    //      - if quue not full, then just add as usual
    // - if minHeap is empty, check the queue whether still contains valid tasks
    pq := NewMinHeap[Task](LessFunc)
    taskCount := make(map[byte]int)
    for i:=0;i<len(tasks);i++{
        taskCount[tasks[i]]++
    }
    queue := make([]Task, 0)
    for taskId, count := range taskCount {
        pq.data = append(pq.data, Task{taskId, count, 0})
    }
    heap.Init(pq)
    
    // seq := make([]byte, 0)
    interval := 0

    for len(queue) > 0 || pq.Len() > 0 {
        interval++
        if pq.Len() > 0 {
            pop := heap.Pop(pq).(Task)
            pop.count--
            // seq = append(seq, pop.id)

            if pop.count > 0 {
                pop.nextTime = interval + n
                queue = append(queue, pop)
            }
        }else{
            // seq = append(seq, idleId)
        }
        if len(queue) > 0 && queue[0].nextTime <= interval {
            heap.Push(pq, queue[0])
            queue = queue[1:]
        }
    }
    // fmt.Println(string(seq))
    return interval

}

var idleTask = Task{idleId, 0, 0}
const idleId = byte('-')

type Task struct {
    id byte
    count int
    nextTime int
}

func (t *Task) IsValid() bool {
    return t.id >= byte('A') && t.id <= byte('Z') && t.count > 0
}

func LessFunc(a, b Task) bool {
    if a.count == b.count {
        return a.id < b.id
    }
    return a.count > b.count
}

type MinHeap[T any] struct {
    data []T
    lessFunc func(a, b T) bool
}

func NewMinHeap[T any](lessFunc func(a, b T) bool, data ...T) *MinHeap[T]{
    res := &MinHeap[T]{
        lessFunc: lessFunc,
    }
    res.data = make([]T, len(data))
    if len(data) > 0 {
        copy(res.data, data)
        heap.Init(res)
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

func assert(cond bool, msg string) {
    if !cond {
        panic(msg)
    }
}