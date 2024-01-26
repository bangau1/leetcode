func hasTrailingZeros(nums []int) bool {
    count := 0

    for _, num := range nums {
        if num % 2 == 0 {
            count++
        }
        if count >= 2 {
            return true
        }
    }

    return false
}