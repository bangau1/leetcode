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
        // fmt.Println(num, count)
        if (count %3) % 2 == 0 {
            total += count/3 + (count %3)/2
            // fmt.Println(num, "total", count/3 + (count %3)/2)
        }else {
            total += count/3 - 1 + (count - 3 * (count/3 - 1)) /2  
            // fmt.Println(num, "total", count/3 - 1 + (count - (count/3 - 1)) /2  )
        }
    }

    return total

}