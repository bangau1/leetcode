func isBipartite(graph [][]int) bool {
    group := make([]byte, len(graph))
    for i:=0;i<len(group);i++{
        group[i] = 255
    }

    // 0 = undecided
    // 1 = red group
    // 2 = green group

    // check first node that has edge
    n := len(graph)
    backlog := make(map[int]bool)
    queue := make([]int, 0)
    for i:=0;i<n;i++{
        if len(graph[i]) > 0 {
            backlog[i] = true
        }
    }
    visited := make([]bool, n)
    for len(backlog) > 0 {
        queue = append(queue, getFirstKey(backlog))    
        for len(queue) > 0 {
            q := queue[0]
            queue = queue[1:]
            delete(backlog, q)
            if visited[q] {
                continue
            }
            visited[q] = true
            if group[q] == 255 {
                group[q] = 0
            }

            for _, next := range graph[q] {
                if group[next] == group[q] {
                    return false
                }
                if visited[next] {
                    continue
                }

                if group[next] == 255 {
                    group[next] = (group[q] + 1) % 2
                    queue = append(queue, next)   
                }else{
                    queue = append(queue, next)   
                }
            }

        }
    }
    fmt.Println(group)
    // for i:=0;i<n;i++{
    //     if group[i] == 255 {
    //         return false
    //     }
    // }
    return true
}

func getFirstKey(data map[int]bool) int {
    for k, _ := range data {
        return k
    }
    return -1
}