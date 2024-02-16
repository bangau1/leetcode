func findLeastNumOfUniqueInts(arr []int, k int) int {
    counter := make(map[int]int)

    for _, num := range arr {
        counter[num]++
    }
    data := make([]int, 0)
    for num, _ := range counter{
        data = append(data, num)
    }

    sort.Slice(data, func(a, b int) bool {
        return counter[data[a]] < counter[data[b]]
    })
    var count int
    var idx int
    for idx < len(data){
        count = counter[data[idx]]
        if k >= count {
            k -= count
        }else{
            break
        }
        idx++
    }
    return len(data)-idx
}