
func longestSubarray(nums []int, limit int) int {
    n := len(nums)
    if n == 1 {
        return 1
    }
    // the idea is to maintain the 2 sliding windows:
    // of minimum element and max element
    // - as long as the max-min <= limit, expand the window
    // - when it break, check the left most max.index and min.index, reset the windows from min(max.index, min.index)+1
    minEls := make([]int, 0)
    maxEls := make([]int, 0)
    res := 1
    idx := 0
    start:=0
    for idx < n {
        for len(maxEls) > 0 && nums[idx] >= nums[maxEls[len(maxEls)-1]] {
            maxEls = maxEls[:len(maxEls)-1]
        }
        for len(minEls) > 0 && nums[idx] <= nums[minEls[len(minEls)-1]] {
            minEls = minEls[:len(minEls)-1]
        }
        
        maxEls = append(maxEls, idx)
        minEls = append(minEls, idx)

        

        if nums[maxEls[0]] - nums[minEls[0]] <= limit {
            res = max(res, idx - start + 1)
            // fmt.Println("start", start)
            // fmt.Println("max", maxEls)
            // fmt.Println("min", minEls)
            idx++
        }else{
            for len(maxEls) > 0 && len(minEls) > 0 && nums[maxEls[0]]-nums[minEls[0]] > limit {
                if maxEls[0] < minEls[0] {
                    start = maxEls[0]+1
                    maxEls = maxEls[1:]
                    
                }else{
                    start = minEls[0]+1
                    minEls = minEls[1:]
                }
            }
            idx++

        }

    }

    return res
}