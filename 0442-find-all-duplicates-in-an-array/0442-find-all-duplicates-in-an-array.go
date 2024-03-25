var counter = make([]int, 100000+1)

func findDuplicates(nums []int) []int {
    arraysFill(counter, 0)
    var res []int
    for _, num := range nums {
        counter[num]++
        if counter[num] > 1 {
            res = append(res, num)
        }
    }
    return res
}

func arraysFill[T any](arr []T, val T) {
    for i:=0;i<len(arr);i++{
        arr[i] = val
    }
}