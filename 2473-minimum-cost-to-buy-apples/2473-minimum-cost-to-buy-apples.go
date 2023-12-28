type minHeap []vertex

func (m minHeap) Less(a, b int) bool {
    return m[a].cost < m[b].cost
}

func (m minHeap) Swap(a, b int) {
    m[a], m[b] = m[b], m[a]
}

func (m minHeap) Len() int{
    return len(m)
}

func (m *minHeap) Push(a any) {
    *m = append(*m, a.(vertex))
}

func (m *minHeap) Pop() any {
    l := m.Len()
    item := (*m)[l-1]
    *m = (*m)[0:l-1]
    return item
}

type vertex struct{
    city, cost int
}

type edge struct {
    city int
    cost int
}

func minCost(n int, roads [][]int, appleCost []int, k int) []int64 {
    // the problem: to answer the minimum cost to buy an apple if we start at city i
    // the complexity: after we buy it, we have to go back to starting city i, but the road cost is pricier with k factor
    // 
    // we can do 2 steps here:
    // 1. find the shortest distance between city i to j. Since we want to generate all city shortest distance, can just use floyd warshall algo (or just do normal djikstra for each starting city)
    // 2. After that we introduce new city from n+1 to 2n. n+i = is the city connected by city i alone with edge = k*shortest distance to i+apple cost at city i
    // So we can use djikstra here to calculate the distance to new city and the first city we find (the minimum cost) is the answer

    globalAdjList := make([][]edge, n)
    for _, road := range roads {
        s, d, c := road[0]-1, road[1]-1, road[2]
        globalAdjList[s] = append(globalAdjList[s], edge{d,c})
        globalAdjList[d] = append(globalAdjList[d], edge{s, c})
    }

    // fmt.Println(dist)

    res := make([]int64, 0)
    for start := 0; start < n;start++{

        dist := djikstraNormal(n, start, globalAdjList)
        adjList := make([][]edge, n)
        // create edge to virtual city
        for i:=0;i<n;i++{
            // from start to city i
            adjList[start] = append(adjList[start], edge{i, dist[i]})

            // the city i to virtual city i
            newCost := dist[i]
            if newCost != math.MaxInt {
                newCost = newCost * (k) + appleCost[i]
            }
            adjList[i] = append(adjList[i], edge{n+i, newCost})
        }
        // fmt.Println("start=", start)
        // fmt.Println(adjList)
        // fmt.Println("--------")
        minCost := djikstraModified(n, start, adjList)
        res = append(res, minCost)
    }
    return res
}


func djikstraNormal(n int, start int, adjList [][]edge) []int {
    pq := make(minHeap, 0)
    visited := make([]bool, n)
    dist := make([]int, n)
    for i:=0;i<n;i++{
        dist[i] = math.MaxInt
    }

    dist[start] = 0
    heap.Push(&pq, vertex{start, 0})
    for pq.Len() > 0 {
        node := heap.Pop(&pq).(vertex)

        if visited[node.city]{
            continue
        }

        visited[node.city] = true

        for _, next := range adjList[node.city]{
            if !visited[next.city] && dist[next.city] > dist[node.city] + next.cost {
                dist[next.city] = dist[node.city] + next.cost
                heap.Push(&pq, vertex{next.city, dist[next.city]})
            }
        }
    }
    return dist
}

func djikstraModified(realN int, start int, adjList [][]edge) int64 {

    n := 2 * realN
    pq := make(minHeap, 0)
    visited := make([]bool, n)
    dist := make([]int, n)
    for i:=0;i<n;i++{
        dist[i] = math.MaxInt
    }

    dist[start] = 0
    heap.Push(&pq, vertex{start, 0})
    for pq.Len() > 0 {
        node := heap.Pop(&pq).(vertex)

        if visited[node.city]{
            continue
        }

        visited[node.city] = true
        if node.city >= realN {
            return int64(node.cost)
        }

        for _, next := range adjList[node.city]{
            if !visited[next.city] && dist[next.city] > dist[node.city] + next.cost {
                dist[next.city] = dist[node.city] + next.cost
                heap.Push(&pq, vertex{next.city, dist[next.city]})
            }
        }
    }
    panic("unexpected")
}

func floydWarshall(n int, adjList [][]edge) [][]int {
    dist := make([][]int, n)
    for i:=0;i<n;i++{
        dist[i] = make([]int, n)
        for j:=0;j<n;j++{
            dist[i][j] = math.MaxInt
            if i == j {
                dist[i][j] = 0
            }
        }
        for _, edge := range adjList[i] {
            dist[i][edge.city] = edge.cost
        }
    }

    for k:=0;k<n;k++{
        for i:=0;i<n;i++{
            for j:=0;j<n;j++{
                // dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
                if dist[i][k] != math.MaxInt && dist[k][j]!=math.MaxInt {
                    dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
                }
            }
        }
    }
    return dist
}