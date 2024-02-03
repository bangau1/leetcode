func minOperations(nums []int) int {
    counter := make(map[int]int)

    for _, num := range nums {
        counter[num]++
    }
    total := 0
    for _, count := range counter {
        if count == 1 {
            return -1
        }
        switch count % 3 {
            case 1 :
                total += count/3-1 + 2
            case 2 :
                total+= count/3 + 1
            default:
                total += count/3
        }

    }

    return total

}