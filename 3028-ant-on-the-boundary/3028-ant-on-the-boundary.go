func returnToBoundaryCount(nums []int) int {
    sum := 0
    total := 0
    for _, num := range nums {
        sum += num
        if sum == 0 {
            total++
        }
    }
    return total
}