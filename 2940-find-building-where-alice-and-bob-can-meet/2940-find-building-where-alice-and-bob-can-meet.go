func leftmostBuildingQueries(h []int, queries [][]int) []int {
    nge := computeNGE(h)
    fmt.Println("nge", nge)
    var ans []int
    for _, query := range queries {
        a, b := query[0], query[1]
        if a == b {
            ans = append(ans, a)
            continue
        }else if a > b {
            a, b = b, a
        }
        
       
        // now a < b
        subAns := -1
        if h[a] < h[b] {
            subAns = b
        }else{
            for b != -1 && h[b] <= h[a]{
                
                b = nge[b]
            }
            if b != -1 && h[b] > h[a] {
                subAns = b
            }
        }
        
        ans = append(ans, subAns)

    }

    return ans
}

func computeNGE(heights []int)[]int {
    n := len(heights)
    stack := make([]node, 0)
    nge := make([]int, n)


    for i:=0;i<n;i++ {
        for len(stack) > 0 && stack[len(stack)-1].val < heights[i] {
            item := stack[len(stack)-1]
            nge[item.idx] = i
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, node{
            heights[i],
            i,
        })
    }
    for len(stack) > 0 {
        item := stack[len(stack)-1]
        nge[item.idx] = -1
        stack = stack[:len(stack)-1]
    }
    return nge
}

type node struct{
    val, idx int
}