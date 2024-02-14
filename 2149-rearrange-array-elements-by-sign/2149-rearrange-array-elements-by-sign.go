func rearrangeArray(nums []int) []int {
    plus := make([]int, 0)
    minus := make([]int, 0)

    for i:=0;i<len(nums);i++{
        if nums[i] > 0 {
            plus = append(plus, nums[i])
        }else{
            minus = append(minus, nums[i])
        }
    }

    for i:=0;i<len(nums);i+=2{
        nums[i] = plus[i/2]
        nums[i+1] = minus[i/2]
    }
    return nums
}