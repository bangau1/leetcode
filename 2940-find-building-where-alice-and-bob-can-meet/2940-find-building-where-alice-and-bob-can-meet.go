type query struct {
    a, b, idx int
}
func leftmostBuildingQueries(h []int, queries [][]int) []int {
    ans := make([]int, len(queries))
    fq := make([]query, 0)
    for i:=0;i<len(queries);i++{
        ans[i] = -1
        if queries[i][0] > queries[i][1]{
            queries[i][0], queries[i][1] = queries[i][1], queries[i][0]
        }
        a, b := queries[i][0], queries[i][1]
        if a == b || h[a] < h[b] {
            ans[i] = b
        }
        if ans[i] == -1 {
            fq = append(fq, query{a, b, i})
        }
    }
    if len(fq) == 0 {
        return ans
    }

    sort.Slice(fq, func(a, b int) bool {
        return fq[a].b > fq[b].b
    })
    // fmt.Println("before", ans)
    // fmt.Println("fq", fq)
    
    stack := make([]int, 0)
    // printStack := func(idxs []int) {
    //     data := make([]int, 0)
    //     for _, idx := range idxs{
    //         data = append(data, h[idx])
    //     }
    //     fmt.Println(data)
    // }

    last := len(h)-1
    for _, q := range fq {
        a, b, ansIdx := q.a, q.b, q.idx
        
        // we want to insert from n-1 to b+1 in monotonic stack in decreasing manner
        for i:=last;i>=b+1;i--{
            for len(stack) > 0 && h[stack[len(stack)-1]] <= h[i] {
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, i)
        }
        // printStack(stack)
        last = b
        // stack is now in decreasing manner, 6,5,4,2
        // we want to look the inside the stack the first item that is <= h[a] let say i, then the answer is at i-1
        idx := sort.Search(len(stack), func(idx int) bool {
            return h[stack[idx]] <= h[a]
        })
        if idx-1 >= 0 {
            ans[ansIdx] = stack[idx-1]
        }
    }
    
    
    return ans
}

// func computeNGE(heights []int)[]int {
//     n := len(heights)
//     stack := make([]node, 0)
//     nge := make([]int, n)


//     for i:=0;i<n;i++ {
//         for len(stack) > 0 && stack[len(stack)-1].val < heights[i] {
//             item := stack[len(stack)-1]
//             nge[item.idx] = i
//             stack = stack[:len(stack)-1]
//         }
//         stack = append(stack, node{
//             heights[i],
//             i,
//         })
//     }
//     for len(stack) > 0 {
//         item := stack[len(stack)-1]
//         nge[item.idx] = -1
//         stack = stack[:len(stack)-1]
//     }
//     return nge
// }

// type node struct{
//     val, idx int
// }