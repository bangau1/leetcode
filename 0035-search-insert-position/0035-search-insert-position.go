func searchInsert(nums []int, target int) int {
    
    l, r := 0, len(nums)-1
    
    if target < nums[l]{
        return l
    }
    if target > nums[r] {
        return r+1
    }
    var m int
    for l <= r {
        m = (r+l)/2
        if target == nums[m] {
            return m
        }else if target > nums[m] {
            l = m +1
        }else{
            r = m-1
        }
    }
    // fmt.Println(m)
    if nums[m] < target {
        return m + 1
    }else if nums[m-1] < target {
        return m
    }
    return m
    
}
