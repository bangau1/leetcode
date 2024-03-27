func countSubarrays(nums []int, k int64) int64 {
    window := make([]int, 0)
    
    var sum int64
    var count int64
    for r:=0;r<len(nums);r++{
        sum += int64(nums[r])
        window = append(window, r)
        for len(window) > 0 && sum * int64(len(window)) >= k {
            sum -= int64(nums[window[0]])
            window = window[1:]
        }
        count += int64(len(window))
        
    }
    return count

}

type data struct {
    index int
    sum int64
}