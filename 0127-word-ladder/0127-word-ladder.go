func ladderLength(beginWord string, endWord string, wordList []string) int {
    if beginWord == endWord {
        return 0
    }

    nStr := len(beginWord)
    // in high level, we need to:
    // 1. to calculate the shortest distance of the endWord to each of word in the list
    // 2. and calculate the shortest distance of the beginWord to each word in the list as well
    // that can be done by using djikstra one source to multiple destinations
    // 
    // To calculate the distance of 2 words, it can be done O(N):
    // - 
    // wordSet pointing the string to the index
    wordSet := make(map[string]int)
    for i, word := range wordList {
        wordSet[word] = i
    }
    if _, ok:=wordSet[beginWord]; !ok {
        wordList = append(wordList, beginWord)
        wordSet[beginWord] = len(wordList)-1
    }
    
    src := wordSet[beginWord]
    if _, ok := wordSet[endWord]; !ok {
        return 0
    }
    target := wordSet[endWord]
    // fmt.Println(wordSet)
    n := len(wordSet)
    if n != len(wordList){
        panic("unexpected")
    }
    adjList := make([][]int, n)
    for i:=0;i<n;i++{
        for j:=i+1;j<n;j++{
            dist := wordDist(wordList[i], wordList[j], nStr)
            if dist == 1 {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
    }
    if len(adjList[target]) == 0 || len(adjList[src]) == 0 {
        return 0
    }

    dist := djikstra(src, adjList)
    
    // for i:=0;i<n;i++{
    //     fmt.Println(i, wordList[i], "->", adjList[i])
    // }
    // fmt.Println(wordList)
    // for i:=0;i<n;i++{
    //     fmt.Println(wordList[src], "->", wordList[i], "=", dist[i])
    // }
    // fmt.Println(src, target, "src, target")
    
    if dist[target] == math.MaxInt {
        return 0
    }
    return dist[target] + 1
}

func djikstra(src int, adjList [][]int) []int {
    n := len(adjList)
    dist := make([]int, n)
    visited := make([]bool, n)
    for i:=0;i<n;i++{
        dist[i] = math.MaxInt
    }
    dist[src] = 0
    pq := NewMinHeap[vertex](func(a, b vertex) bool {
        return a.cost < b.cost
    }, vertex{
        src, 0,
    })

    for pq.Len() > 0 {
        u := heap.Pop(&pq).(vertex)
        if visited[u.node] {
            continue
        }

        visited[u.node] = true
        
        for _, v := range adjList[u.node] {
            if !visited[v] && dist[v] > u.cost + 1 {
                dist[v] = u.cost + 1
                heap.Push(&pq, vertex{
                    v, dist[v],
                })
            }
        }
    }

    return dist

}

type vertex struct {
    node, cost int
}

func copyBytes(a string) [10]byte {
    res := [10]byte{}
    for i:=0;i<len(a);i++{
        res[i] = a[i]
    }
    return res
}

func wordDist(a, b string, n int) int {
    diff := 0
    for i:=0;i<n;i++{
        if a[i] != b[i] {
            diff++
        }
    }

    return diff
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
