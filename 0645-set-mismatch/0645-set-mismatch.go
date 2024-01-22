func findErrorNums(nums []int) []int {
    n := len(nums)
    count := make([]int, n)
    res := make([]int, 2)
    
    for i:=0;i<n;i++{
        count[nums[i]-1]++

        if count[nums[i]-1] == 2{
            res[0] = nums[i]
        }
    }

    for i:=0;i<n;i++{
        if count[i] == 0 {
            res[1] = i+1
            break
        }
    }

    return res
}