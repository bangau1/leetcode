func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)
    n := len(nums)
    prefixSum := make([]int64, n+1)
    for i:=1;i<=n;i++{
        prefixSum[i] = int64(nums[i-1]) + prefixSum[i-1]
    }

    for i:=n-2;i>=0;i--{
        large := int64(nums[i+1])
        if prefixSum[i+1] > large {
            return prefixSum[i+2]
        }
    }
    return -1
}