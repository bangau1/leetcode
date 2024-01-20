func uniqueOccurrences(arr []int) bool {
    
    counting := make([]int, 2001)
    for _, num := range arr {
        idx := num + 1000
        counting[idx]++
    }

    flag := make([]bool, 1001)
    for _, count := range counting {
        if count == 0 {
            continue
        }
        if flag[count]{
            return false
        }
        flag[count] = true
        
    }
    return true
}