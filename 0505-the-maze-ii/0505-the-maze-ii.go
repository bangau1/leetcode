type minHeap []vertex

func (m minHeap) Less(a, b int) bool {
    return m[a].w < m[b].w
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

type vertex struct {
    x, y int
    w int
}

type cell struct {
    x, y int
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
    m, n := len(maze), len(maze[0])
    adjList := make([][][]vertex, m)
    for i:=0;i<m;i++{
        adjList[i] = make([][]vertex, n)

        wall := vertex{i, 0, 0} 
        j:=0 
        for j < n {
            if maze[i][j] == 1 {
                for j < n && maze[i][j] == 1 {
                    j++
                }
            
                if j < n {
                    wall = vertex{i, j, 0}
                }
            }
            if j < n && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], vertex{wall.x, wall.y, weight})
            }
            j++
        }

        wall = vertex{i, n-1, 0} 
        j =n-1
        for j >= 0  {
            if maze[i][j] == 1 {
                for j >=0  && maze[i][j] == 1 {
                    j--
                }
            
                if j >= 0 {
                    wall = vertex{i, j, 0}
                }
            }
            if j >= 0 && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], vertex{wall.x, wall.y, weight})
            }
            j--
        }
    }

    for j:=0;j<n;j++{
        wall := vertex{0, j, 0} 
        i:=0 
        for i < m {
            if maze[i][j] == 1 {
                for i < m && maze[i][j] == 1 {
                    i++
                }
            
                if i < m {
                    wall = vertex{i, j, 0}
                }
            }
            if i < m && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], vertex{wall.x, wall.y, weight})
            }
            i++
        }

        wall = vertex{m-1, j, 0} 
        i =m-1
        for i >= 0 {
            if maze[i][j] == 1 {
                for i >=0 && maze[i][j] == 1 {
                    i--
                }
            
                if i >= 0 {
                    wall = vertex{i, j, 0}
                }
            }
            if i >= 0 && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], vertex{wall.x, wall.y, weight})
            }
            i--
        }
    }
    
    // fmt.Println(adjList)
    inEdge := 0
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            for _, adj := range adjList[i][j] {
                if maze[adj.x][adj.y] == 1 {
                    fmt.Println("found invalid adj(", i, j, ")", adjList[i][j])
                    panic("unexpected")
                }
                if adj.x == destination[0] && adj.y == destination[1] {
                    inEdge+=1
                    break
                }
                
            }
            if inEdge > 0 {
                break
            }
        }
        if inEdge > 0 {
            break
        }
    }
    if inEdge == 0 {
        // fmt.Println("inedge 0")
        return -1
    }

    visited := make([][]bool, m)
    for i:=0;i<m;i++{
        visited[i] = make([]bool, n)
    }
    
    target := vertex{destination[0], destination[1], 0}
    dist := make([][]int, m)
    for i:=0;i<m;i++{
        dist[i] = make([]int, n)
        for j:=0;j<n;j++{
            dist[i][j] = math.MaxInt
        }
    }
    
    
    // do djikstra
    pq := make(minHeap, 0)
    heap.Push(&pq, vertex{start[0], start[1], 0})
    dist[start[0]][start[1]] = 0

    for pq.Len() > 0 {
        node := heap.Pop(&pq).(vertex) 
        if visited[node.x][node.y] {
            continue
        }      

        visited[node.x][node.y] = true
        if node.x == target.x && node.y == target.y {
            return node.w
        }
        if dist[node.x][node.y] == math.MaxInt {
            panic("unexpected")
        }

        for _, v := range adjList[node.x][node.y] {
            if !visited[v.x][v.y] {
                if dist[v.x][v.y] > (dist[node.x][node.y] + v.w){
                    dist[v.x][v.y] = dist[node.x][node.y] + v.w
                }
                heap.Push(&pq, vertex{v.x, v.y, dist[v.x][v.y]})
            }
        }
    }
    if dist[target.x][target.y] == math.MaxInt {
        return -1
    }
    return dist[target.x][target.y]
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}