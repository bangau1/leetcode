func sortedSquares(nums []int) []int {
    n := len(nums)
    idx := sort.SearchInts(nums, 0)
    l, r := idx-1, idx
    var ln, rn int
    res := make([]int, 0)
    for l >= 0 || r < n{
        if l >= 0 && r < n{
            ln, rn = abs(nums[l]), abs(nums[r])
            if ln < rn {
                res = append(res, ln * ln)
                l--
            }else{
                res = append(res, rn * rn)
                r++
            }
            continue
        } 
        
        for l >= 0 {
            res = append(res, nums[l]*nums[l])
            l--
        }

        for r < n  {
            res = append(res, nums[r]*nums[r])
            r++
        }
        
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}