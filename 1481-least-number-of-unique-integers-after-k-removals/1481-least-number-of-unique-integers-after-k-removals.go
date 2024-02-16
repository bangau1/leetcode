func findLeastNumOfUniqueInts(arr []int, k int) int {
    counter := make(map[int]int)

    for _, num := range arr {
        counter[num]++
    }
    occurences := make([]int, len(counter))
    var idx int
    for _, c := range counter{
        occurences[idx] = c
        idx++
    }

    sort.Ints(occurences)
    var count int
    idx = 0
    for idx < len(occurences){
        count = occurences[idx]
        if k >= count {
            k -= count
        }else{
            break
        }
        idx++
    }
    return len(occurences)-idx
}