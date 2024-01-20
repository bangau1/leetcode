func findDuplicate(nums []int) int {
    counting := make([]int, 100000+1)

    for _, num := range nums {
        if counting[num] == 1 {
            return num
        }
        counting[num]++
    }
    return -1
}