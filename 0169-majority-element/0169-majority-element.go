func majorityElement(nums []int) int {
    bound := len(nums)/2
    counter := make(map[int]int)
    for _, num := range nums {
        counter[num]++
    }
    for num, count := range counter {
        if count > bound {
            return num
        }
    }
    return -1
}