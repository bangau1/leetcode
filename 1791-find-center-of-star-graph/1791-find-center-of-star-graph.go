func findCenter(edges [][]int) int {
    n := len(edges)+1
    edgeCount := make([]int, n)

    for _, edge := range edges {
        s, t := edge[0]-1, edge[1]-1
        edgeCount[s]++
        edgeCount[t]++
    }

    for node, count := range edgeCount {
        if count == n -1 {
            return node+1
        }
    }
    return -1
}