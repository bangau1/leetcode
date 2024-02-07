type node string

var dp []int64

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
    costMatrix := shortestDistance(original, changed, cost)
    lenMap := make(map[int]bool)
    
    for _, str := range original {
        lenMap[len(str)] = true
    }
    lenSet := make([]int, 0)
    for length, _ := range lenMap {
        lenSet = append(lenSet, length)
    }

    // fmt.Println(costMatrix)

    dp = make([]int64, len(source)+1)
    for i:=0;i<len(dp);i++{
        dp[i] = math.MaxInt
    }
    dp[0] = 0
    for i:=0;i<len(source);i++{
        if source[i] == target[i] {
            dp[i+1] = 0
        }else{
            break
        }
    }
    // [...i....]
    // dp[i+1] is the cost of changing s[0....i] to t [0....i]
    // dp[i] = is minimum of 
    // a. dp[i-1] + cost[s[i-1]][t[i-1]]
    // b. dp[i-2] + cost[s[i-2..i-1]][t[i-2...i-1]]
    // c. dp[i-max] + cost[s[i-max+1....i-1]][t[i-max+1...i-1]]

    // dp[0] = change empty string, should be 0
    // dp[1] = change 1 character at beginning, dp[0] + cost[s[1-1:1]]
    // dp[2] = chage 2 character at beginning, dp[0] + cost[s[2-2:2]], change 1 char at end-> dp[1] + cost[s[2-1:2]]
    for i:=1;i<len(source)+1;i++{
        if source[i-1] == target[i-1] && dp[i-1] != math.MaxInt {
            dp[i] = min(dp[i], dp[i-1])
        }
        for _, l := range lenSet {
            if i - l < 0 {
                continue
            }
            s := source[i-l:i]
            t := target[i-l:i]
            curr := getCost(costMatrix, s, t)

            if curr != math.MaxInt && dp[i-l] != math.MaxInt {
                dp[i] = min(dp[i], int64(curr) + int64(dp[i-l]))
            }
        }
    }

    // fmt.Println("dp", dp)
    if dp[len(source)] == math.MaxInt64 {
        return -1
    }
    return dp[len(source)]
}

func getCost(costMatrix map[node]map[node]int, source string, target string) int {
    s := node(source)
    t := node(target)
    if s == t {
        return 0
    }

    if _, ok:=costMatrix[s]; !ok {
        return math.MaxInt
    }

    if _, ok:=costMatrix[s][t]; !ok {
        return math.MaxInt
    }
    return costMatrix[s][t]
}



func shortestDistance(original []string, changed []string, cost []int) map[node]map[node]int {
    costMatrix := make(map[node]map[node]int)

    for _, str := range original {
        if _, ok:=costMatrix[node(str)]; !ok {
            costMatrix[node(str)] = make(map[node]int)
        }
    }
    
    for _, str := range changed {
        if _, ok:=costMatrix[node(str)]; !ok {
            costMatrix[node(str)] = make(map[node]int)
        }
    }
    

    for i:=0;i<len(original);i++{
        src, dst := node(original[i]), node(changed[i])
        if _, ok := costMatrix[src][dst]; ok {
            costMatrix[src][dst] = min(costMatrix[src][dst], cost[i])
        }else{
            costMatrix[src][dst] =  cost[i]
        }
    }
    
    for node, _ := range costMatrix {
        costMatrix[node][node] = 0
    }
    // fmt.Println("===========")
    nodes := make([]node,0)
    for node, _ := range costMatrix {
        nodes = append(nodes, node)
        // fmt.Println(node, "=>", costMatrix[node])
    }

    // fmt.Println("===========")
    // n := len(nodes)

    // for inner:=0;inner<n;inner++{
    //     for s:=0;s<n;s++{
    //         for d:=0;d<n;d++{
    //             i := nodes[s]
    //             j := nodes[d]
    //             k := nodes[inner]
    //             _, ok1:=costMatrix[i][k]
    //             _, ok2:=costMatrix[k][j]
    //             if ok1 && ok2 {
    //                 if _, ok:=costMatrix[i][j]; !ok {
    //                     costMatrix[i][j] = costMatrix[i][k] + costMatrix[k][j]
    //                 }else{
    //                     costMatrix[i][j] = min(costMatrix[i][j], costMatrix[i][k] + costMatrix[k][j])
    //                 }
    //             }
    //         }
    //     }
    // }

    for _, source := range nodes {
        res := djikstra(source, costMatrix)
        costMatrix[source] = res

    }

    return costMatrix
}

func djikstra(source node, matrix map[node]map[node]int) map[node]int {
    getCost := func(s, t node) int {
        if _, ok:=matrix[s]; !ok {
            return math.MaxInt
        }
        if _, ok:= matrix[s][t]; !ok {
            return math.MaxInt
        }
        return matrix[s][t]
    }
    pq := NewMinHeap[vertex](func(a, b vertex )bool{
        return a.cost < b.cost
    }, vertex{
        node: source,
        cost: 0,
    })

    visited := make(map[node]bool)

    for pq.Len() > 0 {
        curr := heap.Pop(&pq).(vertex)
        if visited[curr.node] {
            continue
        }

        visited[curr.node] = true
        for next, cost := range matrix[curr.node] {
            if visited[next] {
                continue
            }
            stCost := min(getCost(source, next), curr.cost + cost)
            matrix[source][next] = stCost
            heap.Push(&pq, vertex{
                node: next,
                cost: stCost,
            })
            // distance from source->next
        }

    }

    return matrix[source]
    
}


type MinHeap[T any] struct {
    data []T
    less func(a, b T) bool
}

func NewMinHeap[T any](less func(a, b T)bool, data ...T) MinHeap[T] {
    m := MinHeap[T]{
        data: make([]T, 0),
        less: less,
    }
    if len(data) > 0 {
        for _, datum := range data {
            m.data = append(m.data, datum)
        }
        heap.Init(&m)
    }
    return m
}

func (m MinHeap[T]) Less(a, b int) bool { 
    return m.less(m.data[a], m.data[b])
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



type vertex struct {
    node node
    cost int
}