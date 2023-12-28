type minHeap []weightVertex

func (m minHeap) Less(a, b int) bool {
    if m[a].w == m[b].w {
        return m[a].instruction < m[b].instruction
    }
    return m[a].w < m[b].w
}

func (m minHeap) Swap(a, b int) {
    m[a], m[b] = m[b], m[a]
}

func (m minHeap) Len() int{
    return len(m)
}

func (m *minHeap) Push(a any) {
    *m = append(*m, a.(weightVertex))
}

func (m *minHeap) Pop() any {
    l := m.Len()
    item := (*m)[l-1]
    *m = (*m)[0:l-1]
    return item
}

type weigth struct {
    w int
    i string
}

func (w *weigth) Less(o weigth) bool {
    if w.w == o.w {
        if w.i == "impossible" && o.i == "impossible" {
            return false
        }else if w.i == "impossible" {
            return true
        }else if o.i == "impossible"{
            return false
        }else{
            return w.i < o.i
        }
    }
    return w.w < o.w
}
type weightVertex struct {
    x, y, w int
    instruction string
}

func (w *weightVertex) Cell() cell {
    return cell{w.x, w.y}
}

type cell struct {
    x, y int
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
    m, n := len(maze), len(maze[0])
    adjList := make([][][]weightVertex, m)
    maze[hole[0]][hole[1]] = 2
    for i:=0;i<m;i++{
        adjList[i] = make([][]weightVertex, n)

        wall := cell{i, 0} 
        j:=0 
        for j < n {
            if maze[i][j] == 1 {
                for j < n && maze[i][j] == 1 {
                    j++
                }

            
                if j < n {
                    wall = cell{i, j}
                }
            }else if maze[i][j] == 2 { //ball
                wall = cell{i, j}
            }
            if j < n && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], weightVertex{wall.x, wall.y, weight, "l"})
            }
            j++
        }

        wall = cell{i, n-1} 
        j =n-1
        for j >= 0  {
            if maze[i][j] == 1 {
                for j >=0  && maze[i][j] == 1 {
                    j--
                }
            
                if j >= 0 {
                    wall = cell{i, j}
                }
            }else if maze[i][j] == 2 { //ball
                wall = cell{i, j}
            }
            if j >= 0 && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], weightVertex{wall.x, wall.y, weight, "r"})
            }
            j--
        }
    }

    for j:=0;j<n;j++{
        wall := cell{0, j} 
        i:=0 
        for i < m {
            if maze[i][j] == 1 {
                for i < m && maze[i][j] == 1 {
                    i++
                }
            
                if i < m {
                    wall = cell{i, j}
                }
            }else if maze[i][j] == 2 { //ball
                wall = cell{i, j}
            }
            if i < m && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], weightVertex{wall.x, wall.y, weight, "u"})
            }
            i++
        }

        wall = cell{m-1, j} 
        i =m-1
        for i >= 0 {
            if maze[i][j] == 1 {
                for i >=0 && maze[i][j] == 1 {
                    i--
                }
            
                if i >= 0 {
                    wall = cell{i, j}
                }
            }else if maze[i][j] == 2 { //ball
                wall = cell{i, j}
            }
            if i >= 0 && !(wall.x == i && wall.y == j) {
                weight := abs(i-wall.x) + abs(j-wall.y)
                adjList[i][j] = append(adjList[i][j], weightVertex{wall.x, wall.y, weight, "d"})
            }
            i--
        }
    }
    
    visited := make([][]bool, m)
    for i:=0;i<m;i++{
        visited[i] = make([]bool, n)
    }
    
    target := cell{hole[0], hole[1]}
    dist := make([][]weigth, m)
    for i:=0;i<m;i++{
        dist[i] = make([]weigth, n)
        for j:=0;j<n;j++{
            dist[i][j] = weigth{math.MaxInt, "impossible"}
        }
    }
    
    
    // do djikstra
    pq := make(minHeap, 0)
    heap.Push(&pq, weightVertex{ball[0], ball[1], 0, ""})
    dist[ball[0]][ball[1]] = weigth{0, ""}

    for pq.Len() > 0 {
        node := heap.Pop(&pq).(weightVertex) 
        if visited[node.x][node.y] {
            continue
        }      

        visited[node.x][node.y] = true
        if node.Cell() == target {
            return node.instruction
        }
        if dist[node.x][node.y].w == math.MaxInt {
            panic("unexpected")
        }

        for _, v := range adjList[node.x][node.y] {
            newWeight := weigth{node.w + v.w, node.instruction + v.instruction}
            if !visited[v.x][v.y] && newWeight.Less(dist[v.x][v.y]) {
                dist[v.x][v.y] = newWeight
                heap.Push(&pq, weightVertex{v.x, v.y, dist[v.x][v.y].w, dist[v.x][v.y].i})
            }
        }
    }
    return "impossible"
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}