func canReach(arr []int, start int) bool {
    n := len(arr)
    pq := make([]int, 0)
    visited := make([]bool, n)
    pq = append(pq, start)
    var next int
    for len(pq) > 0 {
        node := pq[0]
        pq = pq[1:]

        if arr[node] == 0 {
            return true
        }
        
        next = node + arr[node]
        if next < n && !visited[next]{
            pq = append(pq, next)
            visited[next] = true
        }

        next = node - arr[node]
        if next >= 0 && !visited[next]{
            pq = append(pq, next)
            visited[next] = true
        }
    }
    return false

}
