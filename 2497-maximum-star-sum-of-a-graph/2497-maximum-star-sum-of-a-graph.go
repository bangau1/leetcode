func maxStarSum(vals []int, edges [][]int, k int) int {
    adjList := make([][]int, len(vals))
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        adjList[a] = append(adjList[a], b)
        adjList[b] = append(adjList[b], a)
    }
    for i:=0;i<len(vals);i++{
        sort.Slice(adjList[i], func(a, b int) bool {
            // sort based on the value decreasing
            return vals[adjList[i][a]] > vals[adjList[i][b]]
        })
    }
    var res = math.MinInt
    for i:=0;i<len(vals);i++{
        subRes := vals[i]
        kc := 0
        for ii:=0;ii<len(adjList[i]);ii++{
            if kc >= k {
                break
            }
            if vals[adjList[i][ii]] < 0 {
                break
            }
            subRes += vals[adjList[i][ii]]
            kc++
            
        }
        if res < subRes {
            res = subRes
        }
    }
    return res
}