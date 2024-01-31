func candy(ratings []int) int {
    n := len(ratings)
    if n == 1 {
        return 1
    }

    count := make([]int, n)
    inEdge := make([]int, n)
    ltAdjList := make([][]int, n)
    for i:=0;i<n;i++{
        if i > 0 && ratings[i] < ratings[i-1] {
            ltAdjList[i] = append(ltAdjList[i], i-1)
            inEdge[i-1]++
        }
        if i < n-1 && ratings[i] < ratings[i+1] {
            ltAdjList[i] = append(ltAdjList[i], i+1)
            inEdge[i+1]++
        }
    }
    q := make([]int, 0)
    for i:=0;i<n;i++{
        if inEdge[i] == 0 {
            q = append(q, i)
        }
    }

// seems like topology sorting?
// a -> b, a need to be taken/decided first, then b can be decided.
// a -> b happened, if a neighbor of b and p[a] < p[b], in|a| = 0
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if count[node] == 0 {
            count[node] = 1
        }
        for _, next := range ltAdjList[node] {
            inEdge[next]--
            count[next] = count[node]+1
            if inEdge[next] == 0 {
                q = append(q, next)
            }
        }
    }
    return sum(count)
}

func sum(data []int) int {
    sum := 0
    for _, datum := range data {
        sum += datum
    }
    return sum
}