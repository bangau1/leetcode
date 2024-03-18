func findMinArrowShots(points [][]int) int {
    // this is quite similar to previous mergeInterval problem
    // given the intervals, we merge it, then return the number of intervals left
    // sort them 
    sort.Slice(points, func(a, b int) bool {
        return points[a][0] < points[b][0]
    })

    res := make([][]int, 0)
    res = append(res, points[0])
    i := 1
    for i < len(points) {
        for i < len(points){
            overlap, newInterval := isOverlap(res[len(res)-1], points[i])
            if !overlap {
                break
            }
            res[len(res)-1][0] = newInterval[0]
            res[len(res)-1][1] = newInterval[1]
            i++
        }

        if i < len(points){
            res = append(res, points[i])
        }
    }

    return len(res)
}

func isOverlap(a, b []int) (bool, [2]int) {
    if a[0] > b[0] {
        a, b = b, a
    }
    // overlap when:
    // 1. a.start <= b.start <= a.end
    overlap := b[0] <= a[1]
    if !overlap {
        return false, [2]int{}
    }

    return overlap, [2]int{b[0], min(a[1], b[1])}
}