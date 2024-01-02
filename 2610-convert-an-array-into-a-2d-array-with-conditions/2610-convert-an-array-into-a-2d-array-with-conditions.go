func findMatrix(nums []int) [][]int {
    data := make(map[int]int)

    for _, num := range nums {
        data[num]++
    }
    res := make([][]int, 0)
    row := 0
    for len(data) > 0 {
        res = append(res, make([]int, 0))
        del := make([]int, 0)
        for num, count := range data {
            res[row] = append(res[row], num)
            data[num] = count-1
            if count == 1 {
                del = append(del, num)
            }
        }

        // delete the key from map
        for _, num := range del {
            delete(data, num)
        }
        row++
    }
    return res

}