type vertex struct {
    x, y int
}

func hasPath(maze [][]int, start []int, destination []int) bool {
    m, n := len(maze), len(maze[0])
    // we can reduce this problem into a graph problem
    // on each cell[i], we need to precompute all its adjacent (the wall) where it can stop.
    // at most, there will be 4 adjacent cells of cell[i], the left, right, bottom, and top wall.
    // 
    // after we already have the adjacent matrix/list, the next thing to do is to find whether we can go 
    // from source to dest
    //
    // We can just use a simple bfs for this
    // Time Complexity: O(M x N)
    // 1. calculate the adjacent list for each cell can be done via dp in 4 direction. O(4 M x N) = O(MN)
    // 2. Then for the BFS: the worst case is we may need to traverse all vertices (considerably linear to M x N). But we can do a simple check, if the inbound edges to destination is zero or outbound from source is zero, then just return false

    adjList := make([][][]vertex, m)
    for i:=0;i<m;i++{
        adjList[i] = make([][]vertex, n)

        wall := vertex{i, 0} 
        j:=0 
        for j < n {
            if maze[i][j] == 1 {
                for j < n && maze[i][j] == 1 {
                    j++
                }
            
                if j < n {
                    wall = vertex{i, j}
                }
            }
            if j < n && !(wall.x == i && wall.y == j) {
                adjList[i][j] = append(adjList[i][j], wall)
            }
            j++
        }

        wall = vertex{i, n-1} 
        j =n-1
        for j >= 0  {
            if maze[i][j] == 1 {
                for j >=0  && maze[i][j] == 1 {
                    j--
                }
            
                if j >= 0 {
                    wall = vertex{i, j}
                }
            }
            if j >= 0 && !(wall.x == i && wall.y == j) {
                adjList[i][j] = append(adjList[i][j], wall)
            }
            j--
        }
    }

    for j:=0;j<n;j++{
        wall := vertex{0, j} 
        i:=0 
        for i < m {
            if maze[i][j] == 1 {
                for i < m && maze[i][j] == 1 {
                    i++
                }
            
                if i < m {
                    wall = vertex{i, j}
                }
            }
            if i < m && !(wall.x == i && wall.y == j) {
                adjList[i][j] = append(adjList[i][j], wall)
            }
            i++
        }

        wall = vertex{m-1, j} 
        i =m-1
        for i >= 0 {
            if maze[i][j] == 1 {
                for i >=0 && maze[i][j] == 1 {
                    i--
                }
            
                if i >= 0 {
                    wall = vertex{i, j}
                }
            }
            if i >= 0 && !(wall.x == i && wall.y == j) {
                adjList[i][j] = append(adjList[i][j], wall)
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
                    // break
                }
                
            }
        }
        if inEdge > 0 {
            break
        }
    }
    if inEdge == 0 {
        // fmt.Println("inedge 0")
        return false
    }

    visited := make([][]bool, m)
    for i:=0;i<m;i++{
        visited[i] = make([]bool, n)
    }
    
    target := vertex{destination[0], destination[1]}
    // do dfs
    stack := make([]vertex, 0)
    stack = append(stack, adjList[start[0]][start[1]]...)
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[0:len(stack)-1]
        // fmt.Println("visit", node, maze[node.x][node.y])
        if maze[node.x][node.y] == 1 {
            panic("unexpected")
        }
        if node == target{
            return true
        }

        for _, v := range adjList[node.x][node.y] {
            if !visited[v.x][v.y] {
                visited[v.x][v.y] = true
                stack = append(stack, v)
            }
        }
    }
    return false
    
}