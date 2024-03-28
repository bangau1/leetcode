func maxSubarrayLength(nums []int, k int) int {
    n := len(nums)
    maxLen := 0

    counter := make(map[int]int)
    l := 0
    for r:=0;r<n;r++{
        counter[nums[r]]++
        for l <= r && counter[nums[r]] > k {
            counter[nums[l]]--
            l++
        } 
        if r-l+1 > maxLen {
            maxLen = r-l+1
        }
    }

    return maxLen
}