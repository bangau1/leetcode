type point struct {
    x, y int
}

func numberOfPairs(inputPoints [][]int) int {
    xSorted := make([]point, 0)

    for _, inputPoint := range inputPoints {
        p := point{inputPoint[0], inputPoint[1]}
        xSorted = append(xSorted, p)
    }

    sort.Slice(xSorted, func(a, b int) bool {
        if xSorted[a].x == xSorted[b].x {
            return xSorted[a].y > xSorted[b].y
        }
        return xSorted[a].x < xSorted[b].x
    })

    // naive solution
    // 0. for each valid pair Chisato(x1, y1) and Takina(x2, y2). Chisato on upperLeft, and Takina lower right. y1 >= y2 && x1 <= x2
    // 1. check whether the pair is sad or not, by checking if there is any other point inside or On the fence.
    n := len(inputPoints)
    totalPairs := 0
    for a:=0;a<n;a++{
        for b:=a+1;b<n;b++{
            chisato := xSorted[a]
            takina := xSorted[b]

            if chisato.x <= takina.x && chisato.y >= takina.y {
                // check whether there is other point inside the rectangle
                // binary search on x1 <= x <= x2 and y2 <= y <= y1
                lower := sort.Search(n, func(idx int) bool {
                    return xSorted[idx].x >= chisato.x
                })

                upper := sort.Search(n, func(idx int) bool {
                    return xSorted[idx].x > takina.x
                })

                if !isSad(chisato, takina, lower, upper, xSorted) {
                    totalPairs++
                    // fmt.Println(chisato, takina)
                }
            }
        }
    }
    return totalPairs
}

func isSad(chisato, takina point, start, endExclusive int, xSorted []point) bool {
    for i:=start;i<endExclusive;i++{
        if xSorted[i] != chisato && xSorted[i] != takina {
            if chisato.x <= xSorted[i].x && xSorted[i].x <= takina.x && 
                takina.y <= xSorted[i].y && xSorted[i].y <= chisato.y {
                    return true
                }
        }
    }
    return false
}