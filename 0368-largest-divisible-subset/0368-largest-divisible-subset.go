func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    n := len(nums)
    prevIndex := make([]int, n)
    dp := make([]int, n)
    for i:=0;i<n;i++{
        prevIndex[i] = i
        dp[i] = 1
    }
    maxSize := dp[0]
    maxIdx := 0
    for i:=1;i<n;i++{    
        for ii:=i-1;ii>=0;ii--{
            if nums[i] % nums[ii] == 0 && dp[i] < 1 + dp[ii]{
                dp[i] = 1 + dp[ii]
                prevIndex[i] = ii
                if maxSize < dp[i] {
                    maxSize = dp[i]
                    maxIdx = i
                }
            } 
        }
    }
    // fmt.Println(dp)
    // fmt.Println(prevIndex)
    res := make([]int, 0)
    curr:= maxIdx
    for true {
        res = append(res, nums[curr])
        
        if curr == prevIndex[curr] {
            break
        }
        curr = prevIndex[curr]
    } 
    return res
}