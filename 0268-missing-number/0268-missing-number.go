func missingNumber(nums []int) int {
    n := len(nums)
    arr := make([]int, 10000)
    for _, num := range nums {
        arr[num] = 1
    }

    for i:=0;i<=n;i++ {
        if arr[i] == 0 {
            return i
        }
    }
    return -1
}