/*
General idea:
- have 2 monotonic dequeu to maintain both minimum and max element in sliding window
- if min and max in windows doesn't hold true:
    - if minEl < minK || maxEl > maxK
        - let mostLeftIdx := min(idx(minEl), idx(maxEl))
        - remove the mostLeftIdx on both until the condition is satisfied (minEL >= minK && maxEl <= maxK)
    - else, it means that minEl >= minK || maxEl <= maxK, then we can keep increasing the window until satisfied.
        - when satified, then the subarray is from [left...right]. 
        - we also need to keep track the lastMinIdx and lastMaxIdx, lastMinIdx = r when nums[r] == minK and lastMaxIdx = r when nums[r] == maxK
        - then let lastHoldMinIdx = min(lastMinIdx, lastMaxIdx), then the position can be [left...lastHoldMinIdx...right], so the total subarray is from lastHoldMinIdx-left+1
        - keep increasing the window and add the subarray total
    
*/
func countSubarrays(nums []int, minK int, maxK int) int64 {
    minQ := NewMonotonicDequeue[int](func(a, b int) bool{
        return nums[a] < nums[b]
    })
    maxQ := NewMonotonicDequeue[int](func(a, b int) bool {
        return nums[a] > nums[b]
    })
    lastMinIdx, lastMaxIdx := -1, -1

    total := int64(0)
    var l int
    n := len(nums)
    var minIdx int
    var lastHoldIdx int

    for r:=0;r<n;r++{
        minQ.PushBack(r)
        maxQ.PushBack(r)

        for minQ.Len() > 0 && maxQ.Len() > 0 &&
             (nums[minQ.Front()] < minK || nums[maxQ.Front()] > maxK) {
            l = min(minQ.Front(), maxQ.Front())
            
            if minQ.Front() <= l {
                minQ.RemoveFront()
            }

            if maxQ.Front() <= l {
                maxQ.RemoveFront()
            }
        }
        if nums[r] == minK {
            lastMinIdx = r
        }
        if nums[r] == maxK {
            lastMaxIdx = r
        }

        if minQ.Len() > 0 && maxQ.Len() > 0 && nums[minQ.Front()] == minK && nums[maxQ.Front()] == maxK {
            minIdx = min(minQ.Front(), maxQ.Front())
            for l < minIdx && (nums[l] > maxK || nums[l] < minK){
                l++
            }
            lastHoldIdx = min(lastMinIdx, lastMaxIdx)
            // fmt.Println(nums[l:r+1])
            total += int64(lastHoldIdx-l+1)
        }
    }

    return total    
}

type MonotonicDequeu[T any] struct{
    data []T
    lessFunc func(a, b T) bool
}

func NewMonotonicDequeue[T any](lessFunc func(a, b T) bool) *MonotonicDequeu[T] {
    return &MonotonicDequeu[T]{
        data: make([]T, 0),
        lessFunc: lessFunc,
    }
}

func (m *MonotonicDequeu[T]) PushBack(val T) {
    for m.Len() > 0 && m.lessFunc(val, m.Back()) {
        m.RemoveBack()
    }
    m.data = append(m.data, val)
}

func (m *MonotonicDequeu[T]) Front() T {
    return m.data[0]
}

func (m *MonotonicDequeu[T]) Back() T {
    return m.data[len(m.data)-1]
}

func (m *MonotonicDequeu[T]) RemoveFront() {
    m.data = m.data[1:]
}

func (m *MonotonicDequeu[T]) RemoveBack() {
    m.data = m.data[:len(m.data)-1]
}

func (m *MonotonicDequeu[T]) Len() int {
    return len(m.data)
}
