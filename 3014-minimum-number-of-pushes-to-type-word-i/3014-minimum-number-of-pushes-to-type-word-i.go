type data struct {
    char rune
    count int
}

func minimumPushes(word string) int {
    count := make(map[rune]int)
    for _, c := range word {
        count[c]++
    }
    
    freq := make([]data, 0)
    
    for char, count := range count{
        freq = append(freq, data{char,count})
    }
    
    sort.Slice(freq, func(a, b int)bool{
        return freq[a].count > freq[b].count
    })
    
    // map of number to the list of letter
    keys := make([]vertex, 0)
    for i:=2;i<=9;i++{
        keys = append(keys, vertex{
            key: i,
            letterPress: make(map[rune]int),
        })
    }
    
    pq := NewMinHeap[vertex](func(a, b vertex)bool{
        return len(a.letterPress) < len(b.letterPress)
    }, keys...)
    
    letterToPress := make(map[rune]int)
    for _, freqData := range freq {
        char := freqData.char
        
        size := len(pq.data[0].letterPress)
        pq.data[0].letterPress[char] = size+1
        heap.Fix(&pq, 0)
        
        letterToPress[char] = size+1
    }
    pressCount:=0
    for _, char := range word {
        pressCount += letterToPress[char]
    }
    
    return pressCount
    
}

type vertex struct {
    key int
    letterPress map[rune]int
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
