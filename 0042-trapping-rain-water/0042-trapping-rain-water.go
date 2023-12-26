type rec struct {
    val, idx int
}

func computeNGE(nums []int, backward bool) []int {
    n := len(nums)
    stack := make([]rec, 0)
    nge := make([]int, n)

    if !backward {
        for r:=0;r<n;r++{
            for len(stack) > 0 && stack[len(stack)-1].val <= nums[r] {
                pop := stack[len(stack)-1]
                nge[pop.idx] = r //point out the nge of item idx is at r index
                stack = stack[0:len(stack)-1] // pop it
            }

            stack = append(stack, rec{nums[r], r})
        }

        for len(stack) > 0 {
            pop := stack[len(stack)-1]
            nge[pop.idx] = -1 // no greater element
            stack = stack[0:len(stack)-1] // pop it
        }

    }else{
        for r:=n-1;r>=0;r--{
            for len(stack) > 0 && stack[len(stack)-1].val <= nums[r] {
                pop := stack[len(stack)-1]
                nge[pop.idx] = r //point out the nge of item idx is at r index
                stack = stack[0:len(stack)-1] // pop it
            }

            stack = append(stack, rec{nums[r], r})
        }

        for len(stack) > 0 {
            pop := stack[len(stack)-1]
            nge[pop.idx] = -1 // no greater element
            stack = stack[0:len(stack)-1] // pop it
        }
    }
    return nge
}
func trap(heights []int) int {
    n := len(heights)
    if n <= 2 {
        return 0 // at least 3 lengths of bar can form the concave
    }
    // the idea is to find the next and prev greater element, so it forms the concave that can trap the water
    nge := computeNGE(heights, false)
    pge := computeNGE(heights, true)

    water := 0
    // fmt.Println(heights)
    // fmt.Println("pge", pge)
    // fmt.Println("nge", pge)
    i:=1
    for i < n {
        subtotal := 0
        
        l := i
        // search the highest in the left, while still maintaining height constraint
        for pge[l] != -1 && nge[l]!=-1{
            l = pge[l]
        }
    
        r := i
        // search the highest in the right, while still maintaining height constraint
        for nge[r] != -1 && pge[r] != -1{
            r = nge[r]
        }

        if i != l && i !=r {
            // fmt.Println("concave", l, i, r)
            h := min(heights[l], heights[r])
            for l+1 < r {
                l++
                diff := h - heights[l]
                if diff >0{
                    subtotal += diff
                }
            }
            i = r+1
        }else{
            i++
        }
        water += subtotal

    }
    return water
}