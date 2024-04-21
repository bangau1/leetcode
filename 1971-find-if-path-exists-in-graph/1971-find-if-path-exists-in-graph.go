func validPath(n int, edges [][]int, source int, destination int) bool {
    visited := make([]bool, n)
    adjList := make([][]int, n)
    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0])
    }

    q := make([]int, 0)
    q = append(q, source)
    var node int
    for len(q) > 0 {
        node = q[0]
        q = q[1:]

        if node == destination {
            return true
        } 
        if visited[node] {
            continue
        }
        visited[node] = true

        for _, next := range adjList[node] {
            if !visited[next] {
                q = append(q, next)
            }
        }
    }
    return false
}