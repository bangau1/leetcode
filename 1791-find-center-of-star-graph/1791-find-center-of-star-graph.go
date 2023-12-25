func findCenter(edges [][]int) int {
    edgeCount := make(map[int]int)

    for _, edge := range edges {
        s, t := edge[0], edge[1]
        edgeCount[s]++
        edgeCount[t]++
    }

    n := len(edgeCount)
    for node, count := range edgeCount {
        if count == n -1 {
            return node
        }
    }
    return -1
}