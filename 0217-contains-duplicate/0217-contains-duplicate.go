func containsDuplicate(nums []int) bool {
    count := make(map[int]int)

    for _, num := range nums {
        count[num]++

        if count[num] >= 2 {
            return true
        }
    }
    return false
}