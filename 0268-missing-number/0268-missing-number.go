func missingNumber(nums []int) int {
    n := len(nums)
    arr := make([]int, n+1)
    for _, num := range nums {
        arr[num] = 1
    }

    for num, count := range arr {
        if count == 0 {
            return num
        }
    }
    return -1
}