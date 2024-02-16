func rotate(nums []int, k int)  {
    c := 0
    n := len(nums)
    i := 0
    k = k % n
    for c < n {
        c += doRotate(i, nums, k)
        i++
    }
    
}

func doRotate(idx int, nums []int, k int) int {
    n := len(nums)
    c := 1
    prev := idx
    next := (prev + k ) % n
    var tmp int
    nums[next], tmp = nums[prev], nums[next]
        
    for next != idx {
        prev = next
        next = (prev + k ) % n
        nums[next], tmp = tmp, nums[next]
        c++
    }
    return c
}